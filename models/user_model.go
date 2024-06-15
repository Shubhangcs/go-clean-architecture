package models

//User Model
type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
}

type UserRepository interface{
	CreateUserTable() error
	RegisterUser(User) error
	LoginUser() error
}