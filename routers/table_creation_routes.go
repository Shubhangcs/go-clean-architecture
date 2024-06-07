package routers

import (
	"github.com/Shubhangcs/go-clean-architecture/controllers"
	"github.com/gorilla/mux"
)
///Queries related to creating table which is a slice of [@string]
var queryStrings = []string{
	"CREATE TABLE FOOD(FOOD_ID INTEGER PRIMARY KEY , FOOD_NAME VARCHAR(30) , FOOD_TYPE VARCHAR(10) , FOOD_PRICE INTEGER);",
	"CREATE TABLE ORDERS(ORDER_ID INTEGER PRIMARY KEY , ORDER_NAME VARCHAR(30) , ORDER_ADDRESS VARCHAR(50) , ORDER_PRICE INTEGER)",
	"CREATE TABLE USER(USER_ID INTEGER PRIMARY KEY , USER_NAME VARCHAR(30) , USER_ADDRESS VARCHAR(50) , USER_PHONE VARCHAR(10))",
}
//Routers related to creating tabels are found here
func InitialTableCreationRoutes(router *mux.Router){
	router.HandleFunc("/create/foodtable" , (&controllers.Queries{QueryString: queryStrings[0]}).CreateTableFunction).Methods("PUT")
	router.HandleFunc("/create/ordertable" , (&controllers.Queries{QueryString: queryStrings[1]}).CreateTableFunction).Methods("PUT")
	router.HandleFunc("/create/usertable" , (&controllers.Queries{QueryString: queryStrings[2]}).CreateTableFunction).Methods("PUT")
}