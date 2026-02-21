package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	dbUser := "root"
	dbPass := "root"
	host := "localhost"
	port := 3303
	dbName := "user_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Fail to Ping dB")
	}
	return db
}
