package controller

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

type Router struct {
	handler CommandHandler
}

func NewRouter(handler CommandHandler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) Handle(ctx context.Context, cmd entity.Command) error {
	if !r.handler.Enabled() {
		return nil
	}

	slog.InfoContext(ctx, r.handler.Name()+": running", "command", cmd.Action)

	err := r.executeAction(ctx, cmd)
	if err != nil {
		return fmt.Errorf("r.executeAction: %w", err)
	}

	return nil
}

func (r *Router) executeAction(ctx context.Context, cmd entity.Command) error {
	switch cmd.Action {
	case entity.CreateAction:
		err := r.handler.Create(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.create: %w", err)
		}
	case entity.GetAction:
		err := r.handler.Get(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.get: %w", err)
		}
	case entity.GetAllAction:
		err := r.handler.GetAll(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.getAll: %w", err)
		}
	case entity.DeleteAction:
		err := r.handler.Delete(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.delete: %w", err)
		}
	default:
		return entity.ErrUnknownCommand
	}

	return nil
}
