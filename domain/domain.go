package domain

import (
	"fmt"
	guuid "github.com/google/uuid"
)

type (
	ID string
)

func (id *ID) String() string {
	return string(*id)
}

func ParseID(id string) (ID, error) {
	guid, err := guuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("id is not valid")
	}

	return ID(guid.String()), nil
}

func NewID() ID {
	guid := guuid.New()

	return ID(guid.String())
}
