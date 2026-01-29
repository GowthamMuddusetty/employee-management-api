package services

import (
	"context"

	"github.com/GowthamMuddusetty/employee-management-api/internal/auth"
	"github.com/GowthamMuddusetty/employee-management-api/internal/config"
	"github.com/GowthamMuddusetty/employee-management-api/internal/models"
	"github.com/GowthamMuddusetty/employee-management-api/internal/repositories"
	"github.com/google/uuid"
)

type AuthService struct {
	repo      *repositories.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: hash,
	}

	return s.repo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return "", err
	}

	return auth.GenerateToken(user.ID, config.Cfg.JWTSecret, config.Cfg.JWTExpiry)
}
