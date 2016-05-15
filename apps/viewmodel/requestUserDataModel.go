package viewmodel

import (
	"bitbucket.org/sofyan_a/linq.im/domains/users"
)

type RequestUserDataModel struct {
	Data  users.User `json:"data"`
	Token string     `json:"token"`
}
