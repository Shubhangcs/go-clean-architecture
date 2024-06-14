package models

//User Model
type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
}