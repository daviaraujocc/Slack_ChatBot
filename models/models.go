package models

import (
	"log"
	"slack-bot/db"
)

type Host struct {
	Id        int
	Host_name string
	Port      string
	Status    string
}

func InsertHost(host, port string) bool {
	db := db.ConnectDB()
	defer db.Close()
	h := findHost(host)
	if (Host{}) != h {
		return false
	} else {

		insertHostSQL := `INSERT INTO Host(host_name, port, status) VALUES (?, ?, ?)`
		stm, err := db.Prepare(insertHostSQL)
		if err != nil {
			log.Println("Erro ao executar statement")
			return false
		}

		_, err = stm.Exec(host, port, "UP")
		if err != nil {
			log.Println("Erro ao executar statement")
			return false
		}

	}
	return true
}

func DeleteHost(host string) bool {
	db := db.ConnectDB()
	defer db.Close()
	deleteHostSQL := `DELETE FROM Host WHERE host_name=?`
	stm, err := db.Prepare(deleteHostSQL)

	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}

	_, err = stm.Exec(host)
	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}
	return true
}

func ResetAllHosts() bool {
	db := db.ConnectDB()
	defer db.Close()
	resetHostSQL := `DELETE FROM Host`
	stm, err := db.Prepare(resetHostSQL)
	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}

	_, err = stm.Exec()
	if err != nil {
		log.Println("Erro ao executar statement")
		return false
	}
	return true
}

func EditStatus(status, host string) {
	db := db.ConnectDB()
	defer db.Close()
	editStatusHostSQL := `UPDATE Host SET status=? WHERE host_name=?`
	stm, _ := db.Prepare(editStatusHostSQL)

	_, _ = stm.Exec(status, host)
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
		err = selectAllHosts.Scan(&h.Id, &h.Host_name, &h.Port, &h.Status)
		if err != nil {
			panic(err.Error())
		}
		hosts = append(hosts, h)
	}
	return hosts
}

func findHost(host_name string) Host {
	db := db.ConnectDB()
	defer db.Close()
	var h Host
	db.QueryRow("SELECT * FROM Host WHERE host_name=?", host_name).Scan(&h.Id, &h.Host_name, &h.Port, &h.Status)
	return h
}
