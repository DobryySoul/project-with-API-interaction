package repo

import (
	"DobryySoul/project-with-API-interaction/internal/entity"
	"DobryySoul/project-with-API-interaction/pkg/logger"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type UserRepository struct {
	Pool   *pgxpool.Pool
	Logger *logger.Logger
}

func NewUserRepository(pg *pgxpool.Pool, logger *logger.Logger) UserRepository {
	return UserRepository{
		Pool:   pg,
		Logger: logger,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.RegisterUser) (int, error) {
	var id int

	query := `INSERT INTO users (username, email, password_hash, created_at) VALUES ($1, $2, $3, $4) RETURNING id;`

	row := r.Pool.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.CreatedAt)
	if err := row.Scan(&user.ID); err != nil {
		r.Logger.L.Info("failed to create user", zap.Error(err))
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1;`
	row := r.Pool.QueryRow(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		r.Logger.L.Info("failed to find user by email", zap.Error(err))
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return &user, nil
}
