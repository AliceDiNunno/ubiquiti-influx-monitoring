package service

import (
	"github.com/AliceDiNunno/gobiquiti"
	"log"
)

func (i *Instance) Tick() {
	err := i.CloudKey.Login()

	if err != nil {
		log.Println("Unable to process tick")
		log.Println(err.Error())
	}

	health, err := i.CloudKey.GetHealth()
	if err != nil {
		log.Println("Unable to fetch health informations")
		log.Println(err.Error())
	}

	clients, err := i.CloudKey.GetClients()
	if err != nil {
		log.Println("Unable to fetch clients informations")
		log.Println(err.Error())
	}

	clientsStats, err := i.CloudKey.GetClientsStats()
	if err != nil {
		log.Println("Unable to fetch clients stats informations")
		log.Println(err.Error())
	}

	clientsMap := map[gobiquiti.Client]gobiquiti.ClientStats{}

	if err == nil {
		for index := range clients.Data {
			client := clients.Data[index]

			for statIndex := range clientsStats.Data {
				stats := clientsStats.Data[statIndex]

				if client.Id == stats.UserId {
					clientsMap[client] = stats
				}
			}
		}
	}

	i.sendHealthMetrics(health.Data)
	i.sendDeviceMetrics(clientsMap)
}
