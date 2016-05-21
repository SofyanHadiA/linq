package viewmodels

import (
	"github.com/SofyanHadiA/linq/domains/products"
)

type RequestProductDataModel struct {
	Data  products.Product `json:"data"`
	Token string           `json:"token"`
}
