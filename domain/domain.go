package domain

import (
	"fmt"
	guuid "github.com/google/uuid"
)

type (
	// ID is a general Unique Identifier that is used across the whole application for each entity
	ID    string
	Error struct {
		OriginError error
		Code        string
		Message     string
	}
)

// String returns string representation of ID
func (id ID) String() string {
	return string(id)
}

// ParseID validates an id and return custom ID if there are no errors
func ParseID(id string) (ID, error) {
	guid, err := guuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("id is not valid")
	}

	return ID(guid.String()), nil
}

// NewID generate a new ID. It uses guid as unique identifier
func NewID() ID {
	guid := guuid.New()

	return ID(guid.String())
}
