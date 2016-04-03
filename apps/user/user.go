package user

type User struct {
	Uid  	  int  		`json:"uid"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	LastLogin int 		`json:"lastLogin"`
}

type Users []User
