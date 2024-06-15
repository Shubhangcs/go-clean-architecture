package repository

import (
	"database/sql"
	"github.com/Shubhangcs/go-clean-architecture/models"
)

type UserRepo struct{
	db *sql.DB
}

func NewUserRepo(sdb *sql.DB) *UserRepo{
	return &UserRepo{
		db: sdb,
	}
}

func (usr *UserRepo) CreateUserTable() error{
	if _ , err := usr.db.Exec("CREATE TABLE USERS(ID INTEGER PRIMARY KEY AUTO_INCREMENT , NAME VARCHAR(30) , EMAIL VARCHAR(30) , PASSWORD VARCHAR(30) , ADDRESS VARCHAR(50))"); err != nil{
		return err;
	}
	return nil
}
func (usr *UserRepo) RegisterUser(u models.User) error{
	if _ , err := usr.db.Exec("INSERT INTO USERS(NAME , EMAIL , PASSWORD , ADDRESS) VALUES(?,?,?,?)",u.Name , u.Email , u.Password , u.Address); err != nil{
		return err;
	}
	return nil
}
func (usr *UserRepo) LoginUser() error{
	if _ , err := usr.db.Exec("INSERT INTO USERS(NAME , EMAIL , PASSWORD , ADDRESS) VALUES(?,?,?,?)"); err != nil{
		return err;
	}
	return nil
}