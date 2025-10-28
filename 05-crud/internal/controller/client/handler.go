package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
)

func (c *Controller) handle(ctx context.Context, cmd *entity.Command) error {
	switch cmd.Action {
	case entity.CreateAction:
		err := c.handleCreate(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.handleCreate: %w", err)
		}
	case entity.GetAction:
		err := c.handleGet(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.handleGet: %w", err)
		}
	case entity.GetAllAction:
		err := c.handleGetAll(ctx)
		if err != nil {
			return fmt.Errorf("c.handleGetAll: %w", err)
		}
	case entity.DeleteAction:
		err := c.handleDelete(ctx, cmd)
		if err != nil {
			return fmt.Errorf("c.delete: %w", err)
		}
	default:
		return entity.ErrUnknownCommand
	}

	return nil
}

func (c *Controller) handleCreate(ctx context.Context, cmd *entity.Command) error {
	name, _ := cmd.Data["name"].(string)
	email, _ := cmd.Data["email"].(string)
	role, _ := cmd.Data["role"].(string)

	dto := dto.CreateUser{
		Name:  name,
		Email: email,
		Role:  role,
	}

	_, err := c.uc.CreateUser(ctx, dto)
	if err != nil {
		return fmt.Errorf("c.uc.CreateUser: %w", err)
	}

	return nil
}

func (c *Controller) handleGet(ctx context.Context, cmd *entity.Command) error {
	str, _ := cmd.Data["uuid"].(string)

	id, err := uuid.Parse(str)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	_, err = c.uc.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("c.uc.GetUser: %w", err)
	}

	return nil
}

func (c *Controller) handleGetAll(ctx context.Context) error {
	_, err := c.uc.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("c.getAll: %w", err)
	}

	return nil
}

func (c *Controller) handleDelete(ctx context.Context, cmd *entity.Command) error {
	str, _ := cmd.Data["uuid"].(string)

	id, err := uuid.Parse(str)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	err = c.uc.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("c.uc.DeleteUser: %w", err)
	}

	return nil
}
