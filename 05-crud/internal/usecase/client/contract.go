package client

import (
	"context"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
)

type UserAdapter interface {
	Create(ctx context.Context, dto dto.CreateUser) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
