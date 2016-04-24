package users

import (
	"database/sql"
	"github.com/satori/go.uuid"
)

type User struct {
	Uid         uuid.UUID      `json:"uid" db:"uid"`
	Avatar      string 		   `json:"photo" db:"avatar"`
	Username    string         `json:"username" db:"username"`
	Email       string         `json:"email" db:"email"`
	Password    string         `json:"password" db:"password"`
	FirstName   string         `json:"firstName" db:"first_name"`
	LastName    string         `json:"lastName" db:"last_name"`
	PhoneNumber string         `json:"phone" db:"phone_number"`
	Address     string         `json:"address" db:"address"`
	Country     string         `json:"country" db:"country"`
	City        string         `json:"city" db:"city"`
	State       string         `json:"state" db:"state"`
	Zip         string         `json:"zip" db:"zip"`
	LastLogin   sql.NullInt64  `json:"lastLogin" db:"last_login"`
	ForgotCode  string         `json:"-" db:"forgot_code"`
	Deleted     string         `json:"-" db:"deleted"`
}

type Users []User

func (user *User) GetId() uuid.UUID {
	return user.Uid
}

func (users *Users) GetLength() int {
	return len(*users)
}
