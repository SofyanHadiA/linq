package repository

import (
	"github.com/satori/go.uuid"
)

type IModel interface {
	GetId() uuid.UUID
}

type IModels interface {
	GetLength() int
}
