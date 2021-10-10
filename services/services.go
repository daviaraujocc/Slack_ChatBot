package services

import (
	"fmt"
	"log"
	"net"
	"slack-bot/messages"
	"slack-bot/models"
	"time"
)

func CheckHosts() {
	n := 0
	for true {
		hosts := models.FindAllHosts()
		for _, host := range hosts {

			_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host.Host_name, host.Port), time.Duration(3)*time.Second)

			if err != nil {
				if host.Status == "UP" {
					models.EditStatus("DOWN", host.Host_name)
					log.Printf("Host %v está DOWN", host.Host_name)
					n++
					data := time.Now()
					bodyText := fmt.Sprintf(":information_source: Notificação #%v \n\n", n) +
						fmt.Sprintf("Host %v está DOWN :x:\n", host.Host_name) +
						fmt.Sprintf("Data: %v \n", data.Format(("01/02/2006"))) +
						fmt.Sprintf("Horário: %v \n\n", data.Format(("15:04:05")))

					messages.MessageSender("notifications", bodyText)
				}
			} else if host.Status == "DOWN" {
				log.Printf("Host %v está UP", host.Host_name)
				models.EditStatus("UP", host.Host_name)
				n++
				data := time.Now()
				bodyText := fmt.Sprintf(":information_source: Notificação #%v \n\n", n) +
					fmt.Sprintf("Host %v está UP :white_check_mark:\n", host.Host_name) +
					fmt.Sprintf("Data: %v \n", data.Format(("01/02/2006"))) +
					fmt.Sprintf("Horário: %v \n\n", data.Format(("15:04:05")))

				messages.MessageSender("notifications", bodyText)
			}

		}
		time.Sleep(10 * time.Second)
	}
}
