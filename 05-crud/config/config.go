package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   AppConfig
		HTTP  HTTPConfig
		Pg    PgConfig
		Redis RedisConfig
		NATS  NATSConfig
		GRPC  GRPCConfig
	}

	AppConfig struct {
		GracefulShutdownTimeout time.Duration `env:"APP_GRACEFUL_SHUTDOWN_TIMEOUT" env-default:"5s"`
	}

	HTTPConfig struct {
		Host              string        `env:"HTTP_HOST"                env-default:"localhost"`
		Port              int           `env:"HTTP_PORT"                env-default:"8080"`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" env-default:"5s"`
	}

	PgConfig struct {
		DSN string `env:"PG_DSN" env-default:"postgres://postgres:postgres@localhost:5432/db05"`
	}

	RedisConfig struct {
		Host       string        `env:"REDIS_HOST"       env-default:"localhost"`
		Port       int           `env:"REDIS_PORT"       env-default:"6379"`
		Expiration time.Duration `env:"REDIS_EXPIRATION" env-default:"5m"`
	}

	NATSConfig struct {
		Enabled bool   `env:"NATS_ENABLED" env-default:"true"`
		URL     string `env:"NATS_URL"     env-default:"nats://nats:password@localhost:4222"`
	}

	GRPCConfig struct {
		Port int `env:"GRPC_PORT" env-default:"50051"`
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
