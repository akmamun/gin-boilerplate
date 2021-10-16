package database

import (
	"database/sql"
	"fmt"
	dbconfig "go-pg/src/config"
)

// dbConnection create database connection
func dbConnection() {
	// connection string and open database
	db, err := sql.Open("postgres", dbconfig.Database())

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
