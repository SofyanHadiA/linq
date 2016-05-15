package viewmodel

import (
	"bitbucket.org/sofyan_a/linq.im/domains/users"
)

type RequestDataUserCredential struct {
	Data  users.UserCredential `json:"data"`
	Token string               `json:"token"`
}
