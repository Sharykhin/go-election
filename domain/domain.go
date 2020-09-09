package domain

import (
	guuid "github.com/google/uuid"
)

type (
	ID string
)

func (id *ID) String() string {
	return string(*id)
}

func NewID() ID {
	id := guuid.New()

	return ID(id.String())
}
