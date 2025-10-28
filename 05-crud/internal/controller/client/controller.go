package client

import (
	"context"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/client"
)

type Controller struct {
	commands []*entity.Command
	uc       *usecase.UserUseCase
}

func New(uc *usecase.UserUseCase, commands []*entity.Command) *Controller {
	router := Controller{
		commands: commands,
		uc:       uc,
	}

	return &router
}

func (c *Controller) Start(ctx context.Context) {
	for _, command := range c.commands {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.handle(ctx, command)
			if err != nil {
				slog.ErrorContext(ctx, "Command execution failed", "error", err)
			}
		}
	}
}
