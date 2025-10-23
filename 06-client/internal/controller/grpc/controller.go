package grpc

import (
	"context"
	"log/slog"

	"github.com/zeuge/hw-go/06-client/config"
	"github.com/zeuge/hw-go/06-client/internal/entity"
	"github.com/zeuge/hw-go/06-client/internal/usecase"
)

type Controller struct {
	enabled bool
	uc      *usecase.UserUseCase
}

func New(cfg *config.GRPCConfig, uc *usecase.UserUseCase) *Controller {
	return &Controller{
		enabled: cfg.Enabled,
		uc:      uc,
	}
}

func (c *Controller) Run(ctx context.Context, cmd entity.Command) error {
	if !c.enabled {
		return nil
	}

	slog.InfoContext(ctx, "GRPC: running", "command", cmd)

	switch cmd.Action {
	case entity.CreateAction:
	case entity.GetAction:
	case entity.GetAllAction:
	case entity.DeleteAction:
	default:
		return entity.ErrUnknownCommand
	}

	return nil
}
