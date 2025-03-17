package service

import (
	"DobryySoul/project-with-API-interaction/internal/entity"
	"DobryySoul/project-with-API-interaction/internal/storage/repo"
	"DobryySoul/project-with-API-interaction/pkg/logger"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository repo.UserRepository
	logger         *logger.Logger
}

func NewService(userRepository repo.UserRepository, logger *logger.Logger) *Service {
	return &Service{
		UserRepository: userRepository,
		logger:         logger,
	}
}

func (s *Service) RegisterUser(ctx context.Context, user *entity.RegisterUser) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.L.Info("failed to hash password", zap.Error(err))
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashedPassword)

	id, err := s.UserRepository.Create(ctx, user)
	if err != nil {
		s.logger.L.Info("failed to create user", zap.Error(err))
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (s *Service) AuthUser(ctx context.Context, email, password string) (*entity.User, error) {
	fmt.Println("email", email, "password", password)
	user, err := s.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		s.logger.L.Info("failed to find user by email", zap.Error(err))
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		s.logger.L.Info("invalid password", zap.Error(err))
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
