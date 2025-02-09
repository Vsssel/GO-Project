package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host = "localhost"
	port = 4444
	user = "assel"
	password = "password"
	dbname = "todo_app"
)


func OpenDatabase () *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)	
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
