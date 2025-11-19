package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/client"
)

type Handler struct {
	uc *usecase.UserUseCase
}

func NewHandler(uc *usecase.UserUseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) Handle(ctx context.Context, commands []*entity.Command) {
	for _, command := range commands {
		select {
		case <-ctx.Done():
			return
		default:
			err := h.handleCommand(ctx, command)
			if err != nil {
				slog.ErrorContext(ctx, "Command execution failed", "error", err)
			}
		}
	}
}

func (h *Handler) handleCommand(ctx context.Context, cmd *entity.Command) error {
	switch cmd.Action {
	case entity.ActionCreate:
		err := h.handleCreate(ctx, cmd)
		if err != nil {
			return fmt.Errorf("h.handleCreate: %w", err)
		}
	case entity.ActionGet:
		err := h.handleGet(ctx, cmd)
		if err != nil {
			return fmt.Errorf("h.handleGet: %w", err)
		}
	case entity.ActionGetAll:
		err := h.handleGetAll(ctx)
		if err != nil {
			return fmt.Errorf("h.handleGetAll: %w", err)
		}
	case entity.ActionDelete:
		err := h.handleDelete(ctx, cmd)
		if err != nil {
			return fmt.Errorf("h.handleDelete: %w", err)
		}
	default:
		return entity.ErrUnknownCommand
	}

	return nil
}

func (h *Handler) handleCreate(ctx context.Context, cmd *entity.Command) error {
	name, _ := cmd.Data["name"].(string)
	email, _ := cmd.Data["email"].(string)
	role, _ := cmd.Data["role"].(string)

	dto := dto.CreateUser{
		Name:  name,
		Email: email,
		Role:  role,
	}

	_, err := h.uc.CreateUser(ctx, dto)
	if err != nil {
		return fmt.Errorf("h.uc.CreateUser: %w", err)
	}

	return nil
}

func (h *Handler) handleGet(ctx context.Context, cmd *entity.Command) error {
	str, _ := cmd.Data["uuid"].(string)

	id, err := uuid.Parse(str)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	_, err = h.uc.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("h.uc.GetUser: %w", err)
	}

	return nil
}

func (h *Handler) handleGetAll(ctx context.Context) error {
	_, err := h.uc.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("h.getAll: %w", err)
	}

	return nil
}

func (h *Handler) handleDelete(ctx context.Context, cmd *entity.Command) error {
	str, _ := cmd.Data["uuid"].(string)

	id, err := uuid.Parse(str)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	err = h.uc.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("h.uc.DeleteUser: %w", err)
	}

	return nil
}
