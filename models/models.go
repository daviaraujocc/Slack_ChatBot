package models

import (
	"database/sql"
	"log"
	"slack-bot/db"
)

type Host struct {
	Id   int
	Host string
	Port string
}

func CreateTableHosts(db *sql.DB) {
	createHostTableSQL := `CREATE TABLE Host (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"host" TEXT,
		"port" TEXT		
	  );`

	log.Println("Create database table...")
	stm, err := db.Prepare(createHostTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stm.Exec()
	log.Println("database table created")
}

func InsertHost(db *sql.DB, host, port string) {
	log.Println("Inserting host record")
	insertHostSQL := `INSERT INTO Host(host, port) VALUES (?, ?)`
	stm, err := db.Prepare(insertHostSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = stm.Exec(host, port)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func FindAllHosts() []Host {
	db := db.ConnectDB()
	defer db.Close()
	selectAllHosts, err := db.Query("SELECT * FROM Host ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	hosts := []Host{}

	for selectAllHosts.Next() {
		h := Host{}
		err = selectAllHosts.Scan(&h.Id, &h.Host, &h.Port)
		if err != nil {
			panic(err.Error())
		}
		hosts = append(hosts, h)
	}
	return hosts
}
