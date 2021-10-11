# About

Hello folks!

This is my 2nd project created while studying for Go language, it's focused on the creation of a BOT for slack platform, with monitoring features (TCP Only) for endpoints and some extras commands.


[![Go Build](https://github.com/DaviAraujoCC/ARIA-ChatBot/actions/workflows/run.yml/badge.svg?branch=main)](https://github.com/DaviAraujoCC/ARIA-ChatBot/actions/workflows/run.yml)[![Go Report Card](https://goreportcard.com/badge/github.com/DaviAraujoCC/Slack_ChatBot)](https://goreportcard.com/report/github.com/DaviAraujoCC/Slack_ChatBot)

## How it works ?

When the application starts, it will create a file based on SQLite named "database.sql" on the same directory of the executable, and it will search for environment variables like BOT_API_TOKEN and some channels ID's that you need to inform in order to work correctly, there is some several routes that will be explained on the Section "Endpoints" that you will use with Slash commands on slack.

## Features:

 - Gorilla Mux ( For HTTP Handler and server )
 - SQLite ( For Embed Database )
 - SlackAPI ( library for Slack API )

## Building the app:

1. Be sure to have `gcc` installed on your O.S because it needs `CGO` since sqlite driver demands it.
2. Run the command: `go build -o slackgo-api` to build the API.
3. Execute and test your bot.
   
##### Or you can use docker:
   `docker build -t your-user/your-app-name .` <br>
   `docker run -t your-user/your-app-name -p 3000:3000(default)`

## Variables:


| Variable | Description |
| --- | --- |
| PORT | Default port for communication with API Server |
| BOT_TOKEN_API | API token of your app created in SLACK |
| MONITOR_CHANNEL | Channel ID of where your commands will be executed |
| NOTIFICATION_CHANNEL | Channel ID of where your notifications will be sended |

## Endpoints:

| Endpoint | Description | Usage |
| --- | --- | --- |
| /api/v1/monitor | Add/Remove endpoints to be monitored | /monitor add host {host} {port} - to add new entrie. <br> /monitor remove host {host} - to remove entrie. |
| /api/v1/hosts | Return all hosts registered in DB | * |
| /api/v1/ping | Ping to target and port informed | /ping {host} {port} |
| /api/v1/help | Show help message | * |
| /api/v1/reset | Reset all entries in DB | * |




