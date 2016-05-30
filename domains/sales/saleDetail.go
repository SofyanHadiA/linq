package sales

import (
	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/domains/products"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type SaleDetail struct {
	core.BasicFields
	SaleId       uuid.UUID        `json:"-" db:"sale_id"`
	ItemId       uuid.UUID        `json:"itemId" db:"item_id"`
	Item         products.Product `json:"item" db:"-"`
	Discount     decimal.Decimal  `json:"discount" db:"discount"`
	DiscountType int              `json:"discountType" db:"discount_type"`
}

type SaleDetails []SaleDetail

func (detail *SaleDetail) GetId() uuid.UUID {
	return detail.Uid
}

func (details *SaleDetails) GetLength() int {
	return len(*details)
}
