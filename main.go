package main

import (
	"log"
	"net/http"

	"github.com/Shubhangcs/go-clean-architecture/routers"
	"github.com/gorilla/mux"
)

// Hello Devs in this golang project we are building a restaurent management system using go
// clean architecture for more info checkout https://github.com/Shubhangcs/go-clean-architecture

func main(){
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

	//Routers
	routers.InitialTableCreationRoutes(router);

	log.Printf("Server is Running at PORT %v" , PORT)

	//Server Configuration 
	log.Fatal(http.ListenAndServe(PORT , router))
	
}