package models

import "io"

//User Model

type UserModel struct{
	UserId string `json:"id"`
	UserName string `json:"name"`
	UserEmail string `json:"email"`
	UserPassword string `json:"password"`
}

type UserModelInterface interface{
	CreateUserTable() error
	RegisterUser(*io.ReadCloser) error
}

