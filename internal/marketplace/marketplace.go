package marketplace

import (
	"context"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/blockchain"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *PostgresRepository
	ethClient *blockchain.ETHClient
}

func NewService(repo *PostgresRepository, ethClient *blockchain.ETHClient) *Service {
	return &Service{repo: repo, ethClient: ethClient}
}

func (s *Service) CreateCollection(ctx context.Context, collection *model.Collection, chainID int64, marketplaceAccountHex string) (*model.Collection, error) {
	collection, err := s.repo.createCollection(ctx, collection)
	if err != nil {
		return nil, fmt.Errorf("failed to create db collection: %w", err)
	}

	contractAddressHex, err := s.ethClient.CreateCollection(ctx, collection, chainID, marketplaceAccountHex)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain collection: %w", err)
	}

	collection.Status = model.PendingTXStatus
	collection.Address = contractAddressHex

	if err := s.repo.updateCollection(ctx, collection); err != nil {
		return nil, fmt.Errorf("failed to update db collection: %w", err)
	}

	return collection, nil
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
