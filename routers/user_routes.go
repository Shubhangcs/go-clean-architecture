package routers

import (
	"database/sql"
	"github.com/Shubhangcs/go-clean-architecture/controllers"
	"github.com/Shubhangcs/go-clean-architecture/repository"
	"github.com/gorilla/mux"
)

func UserRouter(db *sql.DB , router *mux.Router){
	user := repository.NewUserRepo(db)
	cont := controllers.NewHelperInstance(user)

	router.HandleFunc("/user" , cont.CreateTable).Methods("GET")
}