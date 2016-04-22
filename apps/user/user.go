package user

import (
	. "linq/core/repository"
	
    "github.com/satori/go.uuid"
)

type User struct {
	Uid       string    `json:"uid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	LastLogin int    `json:"lastLogin"`
}

type Users []IModel

func (user User) GetId() string {
	return user.Uid
}

func (user *User) InsertVal() []interface{}{	
	result := make([]interface{}, 6)
	
	result[0] = uuid.NewV4()
	result[1] = user.Username
	result[2] = user.Email
	result[3] = user.FirstName
	result[4] = user.LastName
	result[5] = user.Password
	
	return result
}

func (user *User) UpdateVal() []interface{}{
	result := make([]interface{}, 6)
	
	result[5] = user.Uid
	result[0] = user.Username
	result[1] = user.Email
	result[2] = user.FirstName
	result[3] = user.LastName
	result[4] = user.Password

	return result
}