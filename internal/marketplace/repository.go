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
}ç

func (r *PostgresRepository) updateCollection(ctx context.Context, collection *model.Collection) error {
	res, err := r.db.NewUpdate().Model(collection).
		ExcludeColumn("created_at").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil{
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
