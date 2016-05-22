package products

import (
	"time"

	"github.com/satori/go.uuid"
)

type ProductCategory struct {
	Uid         uuid.UUID `json:"uid" db:"uid"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Slug        string    `json:"slug" db:"slug"`
	Deleted     bool      `json:"-" db:"deleted"`
	Created     time.Time `json:"created" db:"created"`
	Updated     time.Time `json:"updated" db:"updated"`
}

type ProductCategories []ProductCategory

func (cat *ProductCategory) GetId() uuid.UUID {
	return cat.Uid
}

func (cats *ProductCategories) GetLength() int {
	return len(*cats)
}
