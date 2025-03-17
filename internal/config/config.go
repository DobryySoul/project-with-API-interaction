package config

import (
	"DobryySoul/project-with-API-interaction/pkg/postgres"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres postgres.Config `yaml:"POSTGRES" env:"POSTGRES"`
	JWTSecret string `yaml:"JWT_SECRET" env:"JWT_SECRET" env-default:"secret"`

	HTTP
}

type HTTP struct {
	Host string `yaml:"HOST" env:"HOST" env-default:"localhost"`
	Port string `yaml:"PORT" env:"PORT" env-default:"8080"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	return &cfg, nil
}
