package controllers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"slack-bot/messages"
	"slack-bot/models"
	"strings"
	"time"
)

//var BOT_TOKEN_API = os.Getenv("BOT_TOKEN_API")

func match(pattern string, value string) bool {
	result, _ := regexp.MatchString(pattern, value)
	return result
}

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t := r.FormValue("text")
	tFormat := strings.Split(t, " ")

	if len(tFormat) <= 2 {
		messages.MessageSender("monitor", ":information_source: [MONITOR] Não identifiquei o que você digitou. Consule /help caso tenha alguma dúvida.")
	} else {
		switch {
		case match(`add host.*`, t) && len(tFormat) == 4:
			host, port := tFormat[2], tFormat[3]
			op := models.InsertHost(host, port)
			if op {
				log.Printf("Adição do Host %v bem sucedida!", host)
				messages.MessageSender("monitor", fmt.Sprintf(":white_check_mark: [MONITOR] Adição do host %v bem sucedida!", host))
			} else {
				log.Printf("Não foi possível adicionar o host %v.", host)
				messages.MessageSender("monitor", fmt.Sprintf(":x: [MONITOR] Não foi possível adicionar o host %v, provavelmente já se encontra cadastrado.", host))
			}

		case match(`remove host.*`, t):
			host := tFormat[2]
			op := models.DeleteHost(host)
			if op {
				log.Printf("Remoção do Host %v bem sucedida!", host)
				messages.MessageSender("monitor", fmt.Sprintf(":white_check_mark: [MONITOR] Adição do host %v bem sucedida!", host))
			} else {
				log.Printf("Não foi possível remover o host %v.", host)
				messages.MessageSender("monitor", fmt.Sprintf(":x: [MONITOR] Não foi possível remove o host %v, verifique se digitou corretamente o nome do mesmo.", host))
			}
		default:
			log.Println("Invalid Route")
			messages.MessageSender("monitor", ":information_source: [MONITOR] Por gentileza, informe todos os parâmetros. Consule /help caso tenha alguma dúvida.")
		}
	}

}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := r.FormValue("text")
	tFormat := strings.Split(t, " ")
	if len(tFormat) <= 1 {
		messages.MessageSender("monitor", ":information_source: [PING] Por gentileza, informe todos os parâmetros. Consule /help caso tenha alguma dúvida.")
	} else {
		host, port := tFormat[0], tFormat[1]
		_, err := net.DialTimeout("tcp", host+":"+port, time.Duration(3)*time.Second)
		if err != nil {
			messages.MessageSender("monitor", fmt.Sprintf(":x: [PING] Não consegui me comunicar com %v", host))
		} else {
			messages.MessageSender("monitor", fmt.Sprintf(":white_check_mark: [PING] Host %v na porta %v está respondendo corretamente", host, port))
		}
	}
}

func HostsHandler(w http.ResponseWriter, r *http.Request) {
	hosts := models.FindAllHosts()
	if len(hosts) == 0 {
		messages.MessageSender("monitor", ":information_source: [HOSTS] Sem hosts cadastrados no momento!")
	} else {
		messages.ShowAllHostsMessage(hosts)
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	models.ResetAllHosts()
	messages.MessageSender("monitor", ":information_source: [RESET] Reset efetuado com sucesso!")
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	helpMessage := "*Comandos:* \n\n" +
		"`/ping` - Realiza ping para o host e porta retornando o status da conexão, exemplo: `/ping contoso.com.br 443` \n\n" +
		"`/monitor` - Usado para adicionar/remover hosts e endpoints do monitoramento. Parâmetros: \n\n" +
		"   * add host {host} {porta} - Adiciona o host no monitoramento, exemplo: `/monitor add host contoso.com.br 443` \n\n" +
		"   * remove host {host} - Remove o host do monitoramento, exemplo: `/monitor remove host contoso.com.br` \n\n" +
		"`/hosts` - Mostra todos os hosts cadastrados no monitoramento. \n\n" +
		"`/help` - Exibe esta mensagem. \n\n"

	messages.MessageSender("monitor", helpMessage)
}
