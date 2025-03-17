package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"POSTGRES_PASS" env:"POSTGRES_PASS" env-default:"2345"`
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"SpotifyDB"`
}

func New(ctx context.Context, config Config) (*pgxpool.Pool, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect potgres: %v", err)
	}

	return conn, nil
}
