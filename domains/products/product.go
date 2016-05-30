package products

import (
	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/datatype"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	core.BasicFields
	Title         string                  `json:"title" db:"title"`
	CategoryTitle string                  `json:"-" db:"cat_title"`
	CategoryId    uuid.UUID               `json:"categoryId" db:"category"`
	Category      ProductCategory         `json:"category" db:"-"`
	BuyPrice      decimal.Decimal         `json:"buyPrice" db:"buy_price"`
	SellPrice     decimal.Decimal         `json:"sellPrice" db:"sell_price"`
	Image         datatype.JsonNullString `json:"image" db:"image"`
	Stock         int                     `json:"stock" db:"stock"`
	Code          string                  `json:"sku" db:"code"`
}

type Products []Product

func (product *Product) GetId() uuid.UUID {
	return product.Uid
}

func (products *Products) GetLength() int {
	return len(*products)
}
