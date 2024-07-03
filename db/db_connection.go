package db

import "database/sql"


type Connection struct{
	DB *sql.DB
}

func GetDatabaseConnectionInstance() *Connection{
	db , err := sql.Open("mysql" , "root:password@tcp(127.0.0.1:3306)/restaurant")

	if err != nil{
		panic(err)
	}

	return &Connection{
		db,
	}
	
}