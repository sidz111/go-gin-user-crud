package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDb() *sql.DB {
	dbUsername := "root"
	dbPassword := "root"
	dbHost := "localhost"
	dbPort := 3303
	dbName := "user_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Fail to Ping dB")
	}
	return db
}
