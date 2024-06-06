package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/Shubhangcs/go-clean-architecture/db"
)

//this controller deals with adding food items and performing operations related to food items

// Database connection instance
var connectToDatabase *db.DBConnection = db.NewDatabaseConnection("root", "password", "restaurant")

// Only Used First Time Just to Create Tables
func CreatingFoodTable(w http.ResponseWriter, r *http.Request) {
	ct := *connectToDatabase
	_, err := ct.ConnectionStr.Exec("CREATE TABLE FOOD(FOOD_ID INTEGER PRIMARY KEY , FOOD_NAME VARCHAR(30) , FOOD_TYPE VARCHAR(10) , FOOD_PRICE INTEGER);")
	if err != nil {
		panic(err)
	}
	defer ct.ConnectionStr.Close()
	//Setting headers to have json data
	w.Header().Set("Content-Type", "application/json")
	//Sending back the response of table creation
	json.NewEncoder(w).Encode("Table Created Successfully")
}
