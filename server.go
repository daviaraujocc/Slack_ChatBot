package main

import (
	"log"
	"net/http"
	"os"
	"slack-bot/controllers"
	"slack-bot/db"
	"slack-bot/models"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	db.CreateDB()
	db := db.ConnectDB()
	models.CreateTableHosts(db)
	defer db.Close()
}

func main() {

	addrPort := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/hello-world", controllers.HelloWorldHandler)
	r.HandleFunc("/api/v1/ping", controllers.PingHandler)
	r.HandleFunc("/api/v1/monitor", controllers.InsertHandler)
	//r.HandleFunc("/api/v1/hosts", controllers.HostsHandler)

	srv := &http.Server{
		Addr:         ":" + addrPort,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)
}
