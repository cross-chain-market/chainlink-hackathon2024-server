package marketplace

import (
	"context"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *PostgresRepository
}

func NewService(repo *PostgresRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCollection(ctx context.Context, collection *model.Collection) (*model.Collection, error) {
	collection, err := s.repo.createCollection(ctx, collection)
	if err != nil {
		return nil, fmt.Errorf("failed to create db collection: %w", err)
	}

	return collection, nil
}

func (s *Service) GetCollection(ctx context.Context, collectionID int64) (*model.Collection, error) {
	return s.repo.getCollection(ctx, collectionID)
}

func (s *Service) GetCollectionByNameAndOwnerAddress(ctx context.Context, collectionName, ownerAddressHes string) (*model.Collection, error) {
	return s.repo.getCollectionByNameAndOwnerAddress(ctx, collectionName, ownerAddressHes)
}

func (s *Service) GetUserCollections(ctx context.Context, ownerAddressHex string) ([]*model.Collection, error) {
	return s.repo.getUserCollections(ctx, ownerAddressHex)
}

func (s *Service) ProcessCollectionDeployed(ctx context.Context, collectionName string, collectionAddress, ownerAddress common.Address) error {
	return s.repo.processCollectionDeployed(ctx, collectionName, collectionAddress.Hex(), ownerAddress.Hex())
}

func (s *Service) ListItem(ctx context.Context, collectionID, itemID, listedAmount int64, fiatPrice float64) (*model.Item, error) {
	item, err := s.repo.listItem(ctx, collectionID, itemID, listedAmount, fiatPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}

	return item, nil
}

func (s *Service) UnlistItem(ctx context.Context, collectionID, itemID, listedAmount int64) (*model.Item, error) {
	item, err := s.repo.unlistItem(ctx, collectionID, itemID, listedAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to unlist items: %w", err)
	}

	return item, nil
}

func (s *Service) GetListings(ctx context.Context, collectionID *int64) ([]*model.Item, error) {
	return s.repo.getListings(ctx, collectionID)
}

func (s *Service) BuyItems(ctx context.Context, collectionID, itemID, amount int64, fromAddress, toAddress string) (*model.Item, error) {
	return s.repo.buyItems(ctx, collectionID, itemID, amount, fromAddress, toAddress)
}

func (s *Service) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return s.repo.registerUser(ctx, user)
}

func (s *Service) LoginUser(ctx context.Context, email, password string) (bool, error) {
	return s.repo.loginUser(ctx, email, password)
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.getUserByEmail(ctx, email)
}
