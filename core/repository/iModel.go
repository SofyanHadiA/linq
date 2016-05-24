package repository

import (
	"time"

	"github.com/satori/go.uuid"
)

type BasicFields struct {
	Uid     uuid.UUID `json:"uid" db:"uid"`
	Deleted bool      `json:"-" db:"deleted"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}

type IModel interface {
	GetId() uuid.UUID
}

type IModels interface {
	GetLength() int
}
