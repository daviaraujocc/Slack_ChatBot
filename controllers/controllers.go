package controllers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"slack-bot/db"

	"github.com/slack-go/slack"
)

var BOT_TOKEN_API = os.Getenv("BOT_TOKEN_API")
var log_enable = false

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

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db := db.ConnectDB()
	defer db.Close()

	t := r.FormValue("text")

	if r.Method == "POST" && match(`add host.*`, t) {
	} else {
		log.Fatal("Invalid route.")
	}

}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	timeout := 3 * time.Second
	_, err := net.DialTimeout("tcp", "www.google.com.br", timeout)
	if err != nil {
		MessageSender(fmt.Sprintf("O host está down %v", err))
	} else {
		MessageSender("O host está UP!")
	}

}

func MessageSender(message string) {
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

func match(pattern string, value string) bool {
	result, _ := regexp.MatchString(pattern, value)
	return result
}
