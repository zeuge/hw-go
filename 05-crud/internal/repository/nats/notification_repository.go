package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/zeuge/hw-go/05-crud/config"
)

type NotificationRepository struct {
	nc      *nats.Conn
	enabled bool
}

func New(cfg *config.NATSConfig) (*NotificationRepository, error) {
	repo := &NotificationRepository{
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

func (r *NotificationRepository) Close() {
	if r.nc != nil {
		r.nc.Close()
	}
}

func (r *NotificationRepository) Publish(subject string, message string) error {
	if !r.enabled {
		return nil
	}

	err := r.nc.Publish(subject, []byte(message))
	if err != nil {
		return fmt.Errorf("r.nc.Publish: %w", err)
	}

	return nil
}
