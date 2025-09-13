package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App AppConfig `mapstructure:"app"`
		Pg  PgConfig  `mapstructure:"pg"`
		Log LogConfig `mapstructure:"log"`
	}

	AppConfig struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}

	PgConfig struct {
		DSN            string `mapstructure:"dsn"`
		MinConnections int    `mapstructure:"min_connections"`
		MaxConnections int    `mapstructure:"max_connections"`
	}

	LogConfig struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	}
)

var defaultConfig = map[string]any{
	"app.host":           "localhost",
	"app.port":           8080,
	"pg.dsn":             "postgres://user:password@localhost:5432/db",
	"pg.min_connections": 2,
	"pg.max_connections": 10,
	"log.level":          "info",
	"log.format":         "text",
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	for key, value := range defaultConfig {
		v.SetDefault(key, value)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	config := new(Config)
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal: %w", err)
	}
	return config, nil
}
