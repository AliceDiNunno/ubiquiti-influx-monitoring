package service

import (
	"fmt"
	"github.com/AliceDiNunno/gobiquiti"
	"github.com/AliceDiNunno/ubiquiti-influx-monitoring/src/infra"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

func NewService(cloudKey gobiquiti.Config, influx infra.InfluxDB) *Instance {
	cloudKeyInstance := gobiquiti.CloudKeyInstance{Config: cloudKey}

	influxClient := influxdb2.NewClient(
		fmt.Sprintf("http://%s:%d/", influx.Hostname, influx.Port),
		fmt.Sprintf("%s:%s", influx.Username, influx.Password))

	return &Instance{
		CloudKey:     cloudKeyInstance,
		influxClient: influxClient,
	}
}

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
