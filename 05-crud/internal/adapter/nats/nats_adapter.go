package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/zeuge/hw-go/05-crud/config"
)

type NATSAdapter struct {
	nc      *nats.Conn
	enabled bool
}

func New(cfg *config.NATSConfig) (*NATSAdapter, error) {
	repo := &NATSAdapter{
		enabled: cfg.Enabled,
	}

	if !cfg.Enabled {
		return repo, nil
	}

	nc, err := nats.Connect(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("nats.Connect: %w", err)
	}

	repo.nc = nc

	return repo, nil
}

func (a *NATSAdapter) Close() {
	if a.nc != nil {
		a.nc.Close()
	}
}

func (a *NATSAdapter) Publish(subject string, message string) error {
	if !a.enabled {
		return nil
	}

	err := a.nc.Publish(subject, []byte(message))
	if err != nil {
		return fmt.Errorf("r.nc.Publish: %w", err)
	}

	return nil
}
