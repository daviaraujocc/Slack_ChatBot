package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
)

var BOT_TOKEN_API = os.Getenv("BOT_TOKEN_API")
var log_enable = false

func main() {

	addrPort := os.Getenv("PORT")
	addr := os.Getenv("HOST")

	r := mux.NewRouter()
	//r.HandleFunc("/commands", CommandsHandler)
	r.HandleFunc("/hello-world", HelloWorldHandler)
	r.HandleFunc("/ping", PingHandler)

	srv := &http.Server{
		Addr:         addr + ":" + addrPort,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	api := slack.New(BOT_TOKEN_API, slack.OptionDebug(log_enable))

	_, _, err := api.PostMessage(
		"C02G50Y5A95", //monitor
		slack.MsgOptionText("hello world", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", "www.google.com.br:443", timeout)
	if err != nil {
		MessageSender(w, r, "O host está UP!")
	} else {
		MessageSender(w, r, fmt.Sprintf("O host está down %v", err))
	}

}

func MessageSender(w http.ResponseWriter, r *http.Request, message string) {
	api := slack.New(BOT_TOKEN_API, slack.OptionDebug(log_enable))

	_, _, err := api.PostMessage(
		"C02G50Y5A95", //monitor
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
