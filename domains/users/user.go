package users

import (
	"time"

	"github.com/SofyanHadiA/linq/core/datatype"

	"github.com/satori/go.uuid"
)

type User struct {
	Uid         uuid.UUID               `json:"uid" db:"uid"`
	Avatar      datatype.JsonNullString `json:"photo" db:"avatar"`
	Username    string                  `json:"username" db:"username"`
	Email       string                  `json:"email" db:"email"`
	Password    string                  `json:"-" db:"password"`
	FirstName   string                  `json:"firstName" db:"first_name"`
	LastName    datatype.JsonNullString `json:"lastName" db:"last_name"`
	PhoneNumber datatype.JsonNullString `json:"phone" db:"phone_number"`
	Address     datatype.JsonNullString `json:"address" db:"address"`
	Country     datatype.JsonNullString `json:"country" db:"country"`
	City        datatype.JsonNullString `json:"city" db:"city"`
	State       datatype.JsonNullString `json:"state" db:"state"`
	Zip         datatype.JsonNullString `json:"zip" db:"zip"`
	LastLogin   datatype.JsonNullInt64  `json:"-" db:"last_login"`
	ForgotCode  datatype.JsonNullString `json:"-" db:"forgot_code"`
	Deleted     bool                    `json:"-" db:"deleted"`
	Created     time.Time               `json:"created" db:"created"`
	Updated     time.Time               `json:"updated" db:"updated"`
}

type Users []User

func (user *User) GetId() uuid.UUID {
	return user.Uid
}

func (users *Users) GetLength() int {
	return len(*users)
}
