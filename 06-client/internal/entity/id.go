package entity

import (
	"fmt"

	"github.com/google/uuid"
)

func NewIDFromString(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("uuid.Parse: %w", err)
	}

	return id, nil
}
