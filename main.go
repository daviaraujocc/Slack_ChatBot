package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
)

var BOT_TOKEN_API = os.Getenv("BOT_TOKEN_API")

func main() {

	addrPort := os.Getenv("PORT")
	addr := os.Getenv("HOST")

	r := mux.NewRouter()
	//r.HandleFunc("/commands", CommandsHandler)
	r.HandleFunc("/hello-world", HelloWorldHandler)

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
	api := slack.New(BOT_TOKEN_API, slack.OptionDebug(true))

	channelId, timestamp, err := api.PostMessage(
		"C02G50Y5A95", //monitor
		slack.MsgOptionText("hello world", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	} else {
		fmt.Printf("A mensagem foi enviada com sucesso %v %v", channelId, timestamp)
	}
}
