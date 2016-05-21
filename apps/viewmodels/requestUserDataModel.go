package viewmodels

import (
	"github.com/SofyanHadiA/linq/domains/users"
)

type RequestUserDataModel struct {
	Data  users.User `json:"data"`
	Token string     `json:"token"`
}
