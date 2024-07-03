package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/Shubhangcs/go-clean-architecture/db"
	"github.com/Shubhangcs/go-clean-architecture/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Hello Devs in this golang project we are building a restaurent management system using go
// clean architecture for more info checkout https://github.com/Shubhangcs/go-clean-architecture

func main(){
	db := db.GetDatabaseConnectionInstance().DB
	router := mux.NewRouter()
	mut := &sync.Mutex{}
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
	PORT := ":9090"
	//Routers
	routers.UserRouters(db , mut , router)

	log.Printf("Server is Running at PORT %v" , PORT)

	//Server Configuration 
	log.Fatal(http.ListenAndServe(PORT , router))
	defer db.Close()
}