package viewmodels

import (
	"github.com/SofyanHadiA/linq/domains/sales"
)

type RequestSaleDataModel struct {
	Data  sales.Sale `json:"data"`
	Token string     `json:"token"`
}
