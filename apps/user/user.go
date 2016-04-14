package user

import (
	. "linq/core/repository"
)

type User struct {
	Uid       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
	result = append(result, user.Uid)
	result = append(result, user.FirstName)
	result = append(result, user.LastName)
	result = append(result, user.Email)
	
	return result
}
