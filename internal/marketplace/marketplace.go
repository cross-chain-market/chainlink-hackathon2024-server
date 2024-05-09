package marketplace

import (
	"context"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
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
		return nil, err
	}

	// TODO: Integrate with Smart Contract

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
