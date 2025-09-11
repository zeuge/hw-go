package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() (ID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return ID{}, fmt.Errorf("uuid.NewV7: %w", err)
	}
	return ID(id), nil
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}
