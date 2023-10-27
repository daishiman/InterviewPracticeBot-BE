package value_objects

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	value uuid.UUID
}

// NewUUID creates a new UUIDValueObject with a generated UUID
func NewUUID() (*UUID, error) {
	return &UUID{value: uuid.New()}, nil
}

func UUIDFromString(s string) (*UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %v", err)
	}
	return &UUID{value: id}, nil
}

func (u *UUID) String() string {
	return u.value.String()
}

func (u *UUID) Equals(other *UUID) bool {
	return u.value == other.value
}
