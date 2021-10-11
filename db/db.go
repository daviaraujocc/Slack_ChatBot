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
	return db
}

func CreateDB() {
	if _, err := os.Stat("./database.db"); os.IsNotExist(err) {

		log.Println("Creating database.db...")

		file, err := os.Create("./database.db")
		if err != nil {
			log.Println("")
		}

		file.Close()
		log.Println("database.db created")
	}
}

func CreateTableHosts() {
	db := ConnectDB()
	defer db.Close()
	createHostTableSQL := `CREATE TABLE Host (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"host_name" TEXT,
		"port" TEXT,
		"status" TEXT	
	  );`

	log.Println("Create database table...")
	stm, err := db.Prepare(createHostTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stm.Exec()
	log.Println("database table created")
}
