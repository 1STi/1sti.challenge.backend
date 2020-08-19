package database

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func ConnectDb(){
	err := godotenv.Load()
	if err != nil { log.Fatal("Error loading .env file") }
  
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")

	// Connect db
	connStr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPass
	Db, err = sql.Open("postgres", connStr)
	if err != nil { log.Fatal(err) }
}