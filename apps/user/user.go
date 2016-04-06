package user

import(
	. "linq/core/repository"
)

type User struct {
	Uid  	  int  		`json:"uid"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	LastLogin int 		`json:"lastLogin"`
}

type Users []IModel

func (user User) SetupModel() IModel{
	return user
}