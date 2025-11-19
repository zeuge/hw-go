package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Scraper ScraperConfig
		File    FileConfig
	}

	ScraperConfig struct {
		Limit   int           `env:"SCRAPER_LIMIT"   env-default:"3"`
		Timeout time.Duration `env:"SCRAPER_TIMEOUT" env-default:"10s"`
		Retries int           `env:"SCRAPER_RETRIES" env-default:"3"`
		Sleep   time.Duration `env:"SCRAPER_SLEEP"   env-default:"500ms"`
	}

	FileConfig struct {
		Input  string `env:"FILE_INPUT"  env-default:"url.txt"`
		Output string `env:"FILE_OUTPUT" env-default:"output.csv"`
	}
)

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}

	return cfg, nil
}
