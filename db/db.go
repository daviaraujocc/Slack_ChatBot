package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
	return db
}

func CreateDB() {
	os.Remove("./database.db")

	log.Println("Creating database.db...")

	file, err := os.Create("./database.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	file.Close()
	log.Println("database.db created")
}
