package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Database connection package

// just an interface to provide higher level of abstraction
type DatabaseConnection interface {
	connectToDatabase() DBConnection
}

// Struct having the ref of database connection
type Connection struct {
	dbUserName string
	dbPassword string
	dbName     string
}

type DBConnection struct {
	ConnectionStr *sql.DB
}

// connect to db function private can only be accessed through the interface
func (c Connection) connectToDatabase() DBConnection {
	con, err := sql.Open("mysql", c.dbUserName+":"+c.dbPassword+"@tcp(127.0.0.1)/"+c.dbName)
	if err != nil {
		panic(err)
	}
	return DBConnection{ConnectionStr: con}
}

// public function to be exported to get access to the database
func NewDatabaseConnection(dbUserName string, dbPassword string, dbName string) *DBConnection {
	var conn DatabaseConnection = Connection{dbUserName: dbUserName, dbPassword: dbPassword, dbName: dbName}
	var realConn = conn.connectToDatabase()
	return &realConn
}
