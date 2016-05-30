package sales

import (
	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/domains/users"

	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Sale struct {
	core.BasicFields
	CustomerId uuid.UUID `json:"customerId" db:"customer"`
	//Customer     customers.Customer `json:"customer" db:"-"`
	UserId       uuid.UUID       `json:"cashierId" db:"user"`
	User         users.User      `json:"cashier" db:"-"`
	Discount     decimal.Decimal `json:"discount" db:"discount"`
	DiscountType int             `json:"discountType" db:"discount_type"`
	Total        decimal.Decimal `json:"total" db:"total"`
	TotalPayment decimal.Decimal `json:"totalPayment" db:"total_payment"`
	PaymentType  int             `json:"paymentType" db:"payment_type"`
	Note         string          `json:"note" db:"note"`
	Detail       SaleDetails     `json:"detail" db:"-"`
}

type Sales []Sale

func (sale *Sale) GetId() uuid.UUID {
	return sale.Uid
}

func (sales *Sales) GetLength() int {
	return len(*sales)
}
