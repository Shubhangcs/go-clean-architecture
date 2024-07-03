package repository

import (
	"database/sql"
	"encoding/json"
	"io"
	"os/exec"
	"sync"

	"github.com/Shubhangcs/go-clean-architecture/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{
	db *sql.DB
	mut *sync.Mutex
}

func NewUserRepository(db *sql.DB , mut *sync.Mutex) *UserRepository {
	return &UserRepository{
		db,
		mut,
	}
}

func (ur *UserRepository) CreateUserTable() error {
	ur.mut.Lock()
	if _ , queryErr := ur.db.Exec("CREATE TABLE USERS(USERID VARCHAR(30) PRIMARY KEY ,USERNAME VARCHAR(30), USEREMAIL VARCHAR(40), USERPASSWORD VARCHAR(30))"); queryErr != nil {
		return queryErr
	}
	defer ur.mut.Unlock()
	return nil
}

func (ur *UserRepository) RegisterUser(r *io.ReadCloser) error {
	ur.mut.Lock()
	var user models.UserModel
	data , readErr := io.ReadAll(*r)
	if readErr != nil {
		return readErr
	}
	if parseErr := json.Unmarshal(data , &user); parseErr != nil {
		return parseErr
	}
	hash , encryptErr := bcrypt.GenerateFromPassword([]byte(user.UserPassword) , 1 )
	if encryptErr != nil {
		return encryptErr
	}
	uid , idErr := exec.Command("uuidgen").Output()
	if idErr != nil {
		return idErr
	}
	if _ , queryErr := ur.db.Exec("INSERT INTO USERS(USERID , USERNAME , USEREMAIL , USERPASSWORD) VALUES(?,?,?,?)" , string(uid) , user.UserName , user.UserEmail , string(hash)); queryErr != nil {
		return queryErr
	}
	defer ur.mut.Unlock()
	return nil
}