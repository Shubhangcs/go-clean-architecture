package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/Shubhangcs/go-clean-architecture/db"
	"github.com/Shubhangcs/go-clean-architecture/models"
)

//this controller deals with adding food items and performing operations related to food items

type Queries struct{
	QueryString string
}

// Database connection instance
var connectToDatabase *db.DBConnection = db.NewDatabaseConnection("root", "password", "restaurant")

func (q *Queries) CreateTableFunction(w http.ResponseWriter , r *http.Request) {
	ct := *connectToDatabase
	_, err := ct.ConnectionStr.Exec(q.QueryString)
	if err != nil {
		panic(err)
	}
	//Closing the database connection and response window at the end of function
	defer ct.ConnectionStr.Close()
	defer r.Body.Close()
	//Setting headers to have json data
	w.Header().Set("Content-Type", "application/json")
	//Sending back the response of table creation
	json.NewEncoder(w).Encode(models.SimplePayload{Message: "Table Created Successfully", Status: http.StatusCreated})
}
