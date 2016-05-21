package viewmodels

import (
	"github.com/SofyanHadiA/linq/domains/users"
)

type RequestDataUserCredential struct {
	Data  users.UserCredential `json:"data"`
	Token string               `json:"token"`
}
