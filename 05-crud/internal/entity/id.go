package entity

import (
	"encoding/json"
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

func NewIDFromString(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ID{}, fmt.Errorf("uuid.Parse: %w", err)
	}

	return ID(id), nil
}

func (id *ID) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(id.String())
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	return data, nil
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	parsed, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	*id = ID(parsed)

	return nil
}

func (id *ID) String() string {
	return uuid.UUID(*id).String()
}
