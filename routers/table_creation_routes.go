package routers

import (
	"github.com/Shubhangcs/go-clean-architecture/controllers"
	"github.com/gorilla/mux"
)


//Routers related to creating tabels are found here
func InitialTableCreationRoutes(router *mux.Router){
	router.HandleFunc("/create/foodtable" , controllers.CreatingFoodTable).Methods("PUT")
}