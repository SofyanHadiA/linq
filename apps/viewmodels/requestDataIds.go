package viewmodels

import (
	"github.com/satori/go.uuid"
)

type data struct {
	Ids []uuid.UUID `json:"ids"`
}

type RequestDataIds struct {
	Data  data   `json:"data"`
	Token string `json:"token"`
}
