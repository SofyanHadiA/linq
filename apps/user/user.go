package user

import (
	. "linq/core/repository"
)

type User struct {
	Uid       int    `json:"uid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	LastLogin int    `json:"lastLogin"`
}

type Users []IModel

func (user User) GetId() int {
	return user.Uid
}

func (user User) InsertVal() []interface{}{
	result := make([]interface{}, 0)
	result = append(result, user.FirstName)
	result = append(result, user.LastName)
	result = append(result, user.Email)
	
	return result
}

func (user User) UpdateVal() []interface{}{
	result := make([]interface{}, 0)
	result = append(result, user.FirstName)
	result = append(result, user.LastName)
	result = append(result, user.Email)
	result = append(result, user.Uid)

	return result
}

var columnMap = map[int]string{
	0 : "uid",
	1 : "username",
}