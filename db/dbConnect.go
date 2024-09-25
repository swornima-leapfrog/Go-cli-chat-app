package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func DbConnect() *pg.DB {
	err:= godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}


	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"), // PostgreSQL server address
		User:     os.Getenv("DB_USER"),  // PostgreSQL username
		Password: os.Getenv("DB_PASSWORD"),  // PostgreSQL password
		Database: os.Getenv("DB_NAME"),    // PostgreSQL database name
	})

	// Check if the connection is successful
	var version string

	_, err = db.QueryOne(pg.Scan(&version), "SELECT version()")

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL")
	}

	fmt.Print("Connected to PostgreSQL")
	return db
}
