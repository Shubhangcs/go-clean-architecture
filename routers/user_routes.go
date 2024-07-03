package routers

import (
	"database/sql"
	"sync"

	"github.com/Shubhangcs/go-clean-architecture/controllers"
	"github.com/Shubhangcs/go-clean-architecture/repository"
	"github.com/gorilla/mux"
)


func UserRouters(db *sql.DB , mut *sync.Mutex , router *mux.Router){
	repo := repository.NewUserRepository(db , mut)
	cont := controllers.NewUserController(repo)

	router.HandleFunc("/createuser" , cont.CreateUserTableRequest).Methods("GET")
	router.HandleFunc("/register" , cont.RegisterUserRequest).Methods("POST")
}