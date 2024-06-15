package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/Shubhangcs/go-clean-architecture/routers"
	"github.com/gorilla/mux"
)

// Hello Devs in this golang project we are building a restaurent management system using go
// clean architecture for more info checkout https://github.com/Shubhangcs/go-clean-architecture

func main(){
	db , _ := sql.Open("mysql" , "root:password@tcp(127.0.0.1:3306)/restaurant")
	//packages being used
	/*
	-> fmt
	-> log
	-> os
	-> sync
	-> net/http
	-> gorilla mux
	-> io
	-> mysql
	*/
	router := mux.NewRouter()
	PORT := ":8000"
	routers.UserRouter(db , router)

	//Routers
	// routers.InitialTableCreationRoutes(router);

	log.Printf("Server is Running at PORT %v" , PORT)

	//Server Configuration 
	log.Fatal(http.ListenAndServe(PORT , router))
	
}