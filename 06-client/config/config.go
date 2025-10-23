package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP HTTPConfig
		GRPC GRPCConfig
	}

	HTTPConfig struct {
		Enabled bool   `env:"HTTP_ENABLED"     env-default:"false"`
		Host    string `env:"HTTP_SERVER_HOST" env-default:"localhost"`
		Port    int    `env:"HTTP_SERVER_PORT" env-default:"8080"`
	}

	GRPCConfig struct {
		Enabled bool   `env:"GRPC_ENABLED"     env-default:"true"`
		Host    string `env:"GRPC_SERVER_HOST" env-default:"localhost"`
		Port    int    `env:"GRPC_SERVER_PORT" env-default:"50051"`
	}
)

func ReadConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}

	return &cfg, nil
}
