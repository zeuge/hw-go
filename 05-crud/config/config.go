package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App        AppConfig
		HTTPServer HTTPServerConfig
		HTTPClient HTTPClientConfig
		Pg         PgConfig
		Redis      RedisConfig
		NATS       NATSConfig
		GRPCServer GRPCServerConfig
		GRPCClient GRPCClientConfig
		Tracing    TracingConfig
	}

	AppConfig struct {
		AppName                 string        `env:"APP_NAME"                      env-default:"app"`
		GracefulShutdownTimeout time.Duration `env:"APP_GRACEFUL_SHUTDOWN_TIMEOUT" env-default:"5s"`
		UseGRPC                 bool          `env:"APP_USE_GRPC"                  env-default:"false"`
	}

	HTTPServerConfig struct {
		Host              string        `env:"HTTP_SERVER_HOST"         env-default:"localhost"`
		Port              int           `env:"HTTP_SERVER_PORT"         env-default:"8080"`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" env-default:"5s"`
		UsePprof          bool          `env:"APP_USE_PPROF"            env-default:"true"`
	}

	HTTPClientConfig struct {
		Host string `env:"HTTP_CLIENT_HOST" env-default:"localhost"`
		Port int    `env:"HTTP_CLIENT_PORT" env-default:"8080"`
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

	GRPCServerConfig struct {
		Port int `env:"GRPC_SERVER_PORT" env-default:"50051"`
	}

	GRPCClientConfig struct {
		Host string `env:"GRPC_CLIENT_HOST" env-default:"localhost"`
		Port int    `env:"GRPC_CLIENT_PORT" env-default:"50051"`
	}

	TracingConfig struct {
		AppName  string `env:"APP_NAME"        env-default:"app"`
		Endpoint string `env:"JAEGER_ENDPOINT" env-default:"localhost:4318"`
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
