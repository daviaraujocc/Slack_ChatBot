package messages

import (
	"fmt"
	"log"
	"os"
	"slack-bot/models"
	"strings"

	"github.com/slack-go/slack"
)

var BOT_TOKEN_API = os.Getenv("BOT_TOKEN_API")

var MONITOR_CHANNEL = os.Getenv("MONITOR_CHANNEL")

var NOTIFICATION_CHANNEL = os.Getenv("NOTIFICATION_CHANNEL")

var log_enable = false

func MessageSender(dst, message string) {
	api := slack.New(BOT_TOKEN_API, slack.OptionDebug(log_enable))
	if dst == "notification" {
		dst = NOTIFICATION_CHANNEL
	} else if dst == "monitor" {
		dst = MONITOR_CHANNEL
	}
	_, _, err := api.PostMessage(
		dst, //monitor
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func ShowAllHostsMessage(hosts []models.Host) {
	api := slack.New(BOT_TOKEN_API, slack.OptionDebug(log_enable))
	var bodyText strings.Builder
	var status string

	for i, host := range hosts {
		if host.Status == "DOWN" {
			status = ":x:"
		} else {
			status = ":white_check_mark:"
		}
		bodyText.WriteString(fmt.Sprintf("ID: %v \n", i+1) +
			fmt.Sprintf("Host: %v \n", host.Host_name) +
			fmt.Sprintf("Porta: %v \n", host.Port) +
			fmt.Sprintf("Status: %v \n\n", status))

	}
	preText := "*Lista de hosts cadastrados:*"
	preTextF := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	bodyTextF := slack.NewTextBlockObject("mrkdwn", bodyText.String(), false, false)

	dividerSection := slack.NewDividerBlock()
	preTextSection := slack.NewSectionBlock(preTextF, nil, nil)
	bodyTextSection := slack.NewSectionBlock(bodyTextF, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection,
		bodyTextSection,
	)

	_, _, _, err := api.SendMessage(
		MONITOR_CHANNEL, msg,
	)
	if err != nil {
		log.Println("error")
	}
}
