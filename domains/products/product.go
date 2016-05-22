package products

import (
	"time"

	"github.com/SofyanHadiA/linq/core/datatype"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Uid        uuid.UUID               `json:"uid" db:"uid"`
	Title      string                  `json:"title" db:"title"`
	BuyPrice   decimal.Decimal         `json:"buyPrice" db:"buy_price"`
	SellPrice  decimal.Decimal         `json:"sellPrice" db:"sell_price"`
	Image      datatype.JsonNullString `json:"image" db:"image"`
	Stock      int                     `json:"stock" db:"stock"`
	Code       string                  `json:"sku" db:"code"`
	Deleted    bool                    `json:"-" db:"deleted"`
	Created    time.Time               `json:"created" db:"created"`
	Updated    time.Time               `json:"updated" db:"updated"`
	CategoryId uuid.UUID               `json:"categoryId" db:"category"`
	Category   ProductCategory         `json:"category" db:"-"`
}

type Products []Product

func (product *Product) GetId() uuid.UUID {
	return product.Uid
}

func (products *Products) GetLength() int {
	return len(*products)
}
