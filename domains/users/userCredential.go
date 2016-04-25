package users

import (
	"github.com/satori/go.uuid"
)

type UserCredential struct {
	Uid         uuid.UUID      `json:"uid" db:"uid"`
	Password      string 		   `json:"password" db:"password"`
	NewPassword       string         `json:"newPassword" db:"-"`
	PasswordConfirm    string         `json:"passwordConfirm" db:"-"`
	RecoveryCode    string         `json:"recoveryCode" db:"recodery_code"`
}

func (user *UserCredential) GetId() uuid.UUID {
	return user.Uid
}