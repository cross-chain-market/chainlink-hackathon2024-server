package marketplace

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	errs "github.com/cross-chain-market/chainlink-hackathon2024-server/internal/errors"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type PostgresRepository struct {
	db *bun.DB
}

func NewPostgresRepository(db *bun.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) createCollection(ctx context.Context, collection *model.Collection) (*model.Collection, error) {
	if err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(collection).Exec(ctx); err != nil {
			return fmt.Errorf("failed to insert collection: %w", err)
		}

		for i := range collection.Items {
			collection.Items[i].CollectionID = collection.ID
		}

		if _, err := tx.NewInsert().Model(&collection.Items).Exec(ctx); err != nil {
			return fmt.Errorf("failed to insert items: %w", err)
		}

		return nil
	}); err != nil {
		var pgErr pgdriver.Error
		ok := errors.As(err, &pgErr)

		if ok && pgErr.IntegrityViolation() && isUniqueViolation(pgErr) {
			return nil, errs.ErrDuplicatedEntity
		}
		return nil, fmt.Errorf("failed to run in transaction: %w", err)
	}

	return collection, nil
}

func (r *PostgresRepository) getCollection(ctx context.Context, collectionID int64) (*model.Collection, error) {
	collection := new(model.Collection)

	err := r.db.NewSelect().
		Model(collection).
		Relation("Items").
		Where("id = ?", collectionID).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrEntityNotFound
		}

		return nil, fmt.Errorf("failed to get collection: %w", err)
	}

	return collection, nil
}

func (r *PostgresRepository) getCollectionByNameAndOwnerAddress(ctx context.Context, collectionName, ownerAddressHex string) (*model.Collection, error) {
	collection := new(model.Collection)

	err := r.db.NewSelect().
		Model(collection).
		Relation("Items").
		Where("name = ?", collectionName).
		Where("owner_address_hex = ?", ownerAddressHex).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrEntityNotFound
		}

		return nil, fmt.Errorf("failed to get collection: %w", err)
	}

	return collection, nil
}

func (r *PostgresRepository) processCollectionDeployed(ctx context.Context, collectionName, collectionAddress, ownerAddressHex string) error {
	if err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		collection := new(model.Collection)

		err := tx.NewSelect().
			Model(collection).
			Where("name = ?", collectionName).
			Where("owner_address_hex = ?", ownerAddressHex).
			Scan(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.ErrEntityNotFound
			}

			return fmt.Errorf("failed to get collection: %w", err)
		}

		collection.Address = collectionAddress
		collection.Status = model.DeployedStatus

		res, err := tx.NewUpdate().Model(collection).
			ExcludeColumn("created_at").
			WherePK().
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("error getting rows affected: %w", err)
		}

		if rowsAffected == 0 {
			return errs.ErrEntityNotFound
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to run in transaction: %w", err)
	}

	return nil
}

func (r *PostgresRepository) getUserCollections(ctx context.Context, ownerAddressHex string) ([]*model.Collection, error) {
	var collections []*model.Collection

	err := r.db.NewSelect().
		Model(&collections).
		Relation("Items").
		Where("owner_address_hex = ?", ownerAddressHex).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user collections: %w", err)
	}

	return collections, nil
}

func (r *PostgresRepository) updateCollection(ctx context.Context, collection *model.Collection) error {
	res, err := r.db.NewUpdate().Model(collection).
		ExcludeColumn("created_at").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errs.ErrEntityNotFound
	}

	return nil
}

func (r *PostgresRepository) listItem(ctx context.Context, collectionID, itemID, listedAmount int64, fiatPrice float64) (*model.Item, error) {
	item := new(model.Item)

	if err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		err := tx.NewSelect().
			Model(item).
			Where("id = ?", itemID).
			Where("collection_id = ?", collectionID).
			Scan(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.ErrEntityNotFound
			}

			return fmt.Errorf("failed to get item: %w", err)
		}

		// if already listed amount + new listed amount > total available amount, then we fail
		if item.ListedAmount+listedAmount > item.TotalAmount {
			return errs.ErrListedAmountGreaterThanTotalAmount
		}

		// if there is some fiat price and the new fiat price is different than the existing one, we fail
		// TODO: This is a valid case but I think we should leave this out of MVP for now since we would need to update blockchain
		if item.FiatPrice != 0 && item.FiatPrice != fiatPrice {
			return errs.ErrCannotUpdateFiatPrice
		}

		item.FiatPrice = fiatPrice
		item.ListedAmount += listedAmount

		_, err = tx.NewUpdate().Model(item).
			ExcludeColumn("created_at").
			WherePK().
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to run in transaction: %w", err)
	}

	return item, nil
}

func (r *PostgresRepository) unlistItem(ctx context.Context, collectionID, itemID, listedAmount int64) (*model.Item, error) {
	item := new(model.Item)

	if err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		err := tx.NewSelect().
			Model(item).
			Where("id = ?", itemID).
			Where("collection_id = ?", collectionID).
			Scan(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.ErrEntityNotFound
			}

			return fmt.Errorf("failed to get item: %w", err)
		}

		if listedAmount > item.ListedAmount {
			return errs.ErrCannotUnlistGreaterAmountThanListedAmount
		}

		item.ListedAmount -= listedAmount

		_, err = tx.NewUpdate().Model(item).
			ExcludeColumn("created_at").
			WherePK().
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to run in transaction: %w", err)
	}

	return item, nil
}

func (r *PostgresRepository) getListings(ctx context.Context, collectionID *int64) ([]*model.Item, error) {
	var items []*model.Item

	query := r.db.NewSelect().Model(&items).Where("listed_amount > 0")

	if collectionID != nil {
		query = query.Where("collection_id = ?", collectionID)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, fmt.Errorf("failed to fetch listings: %w", err)
	}

	// TODO: Improve
	result := make([]*model.Item, 0, len(items))
	deployedCollections := make(map[int64]bool)
	for _, item := range items {
		isDeployed, exists := deployedCollections[item.CollectionID]
		if exists {
			if isDeployed {
				result = append(result, item)
			}

			continue
		}

		ok, err := r.db.NewSelect().
			Model((*model.Collection)(nil)).
			Where("id = ?", item.CollectionID).
			Where("status = ?", model.DeployedStatus).
			Exists(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to check collection status: %w", err)
		}

		if ok {
			result = append(result, item)
			deployedCollections[item.CollectionID] = true
		} else {
			deployedCollections[item.CollectionID] = false
		}
	}

	return result, nil
}

func (r *PostgresRepository) buyItems(ctx context.Context, collectionID, itemID, amount int64, fromAddress, toAddress string) (*model.Item, error) {
	item := new(model.Item)

	if err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		err := tx.NewSelect().
			Model(item).
			Where("id = ?", itemID).
			Where("collection_id = ?", collectionID).
			Scan(ctx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errs.ErrEntityNotFound
			}

			return fmt.Errorf("failed to get item: %w", err)
		}

		// check if amount is valid
		if amount > item.ListedAmount || amount > item.TotalAmount {
			return errs.ErrCannotBuyMoreThanListedAmount
		}

		item.ListedAmount -= amount
		item.TotalAmount -= amount

		_, err = tx.NewUpdate().Model(item).
			ExcludeColumn("created_at").
			WherePK().
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to buy items: %w", err)
	}

	return item, nil
}

func (r *PostgresRepository) registerUser(ctx context.Context, user *model.User) (*model.User, error) {
	if _, err := r.db.NewInsert().Model(user).Exec(ctx); err != nil {
		var pgErr pgdriver.Error
		ok := errors.As(err, &pgErr)

		if ok && pgErr.IntegrityViolation() && isUniqueViolation(pgErr) {
			return nil, errs.ErrDuplicatedEntity
		}

		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return user, nil
}

func (r *PostgresRepository) loginUser(ctx context.Context, email, password string) (bool, error) {
	u, err := r.getUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrEntityNotFound) {
			return false, nil
		}

		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (r *PostgresRepository) getUserByEmail(ctx context.Context, email string) (*model.User, error) {
	u := new(model.User)

	err := r.db.NewSelect().
		Model(u).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrEntityNotFound
		}

		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return u, nil
}

func isUniqueViolation(pgErr pgdriver.Error) bool {
	return strings.Contains(pgErr.Error(), "23505")
}
