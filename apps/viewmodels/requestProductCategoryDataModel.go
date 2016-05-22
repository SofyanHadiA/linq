package viewmodels

import (
	"github.com/SofyanHadiA/linq/domains/products"
)

type RequestProductCategoryDataModel struct {
	Data  products.ProductCategory `json:"data"`
	Token string                   `json:"token"`
}
