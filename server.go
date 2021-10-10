package main

import (
	"log"
	"net/http"
	"os"
	"slack-bot/controllers"
	"slack-bot/db"
	"slack-bot/services"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	db.CreateDB()
	db.CreateTableHosts()
}

func main() {

	addrPort := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ping", controllers.PingHandler)
	r.HandleFunc("/api/v1/monitor", controllers.MonitorHandler)
	r.HandleFunc("/api/v1/hosts", controllers.HostsHandler)
	r.HandleFunc("/api/v1/reset", controllers.ResetHandler)
	r.HandleFunc("/api/v1/help", controllers.HelpHandler)

	if len(addrPort) == 0 {
		addrPort = "30000"
	}

	srv := &http.Server{
		Addr:         ":" + addrPort,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//Daemons
	go services.CheckHosts()

	err := srv.ListenAndServe()
	log.Fatal(err)

}
