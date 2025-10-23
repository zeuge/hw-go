package controller

import (
	"context"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

type CommandHandler interface {
	Enabled() bool
	Name() string
	Create(ctx context.Context, cmd entity.Command) error
	Get(ctx context.Context, cmd entity.Command) error
	GetAll(ctx context.Context, cmd entity.Command) error
	Delete(ctx context.Context, cmd entity.Command) error
}
