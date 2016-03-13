package user

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	LastLogin time.Time `json:"lastLogin"`
}

type Users []User
