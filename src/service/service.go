package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"time"
)

var cloudKeyInformations infra.UbiquitiServer
var influxDB infra.InfluxDB

func LoadService(cloudKey infra.UbiquitiServer, influx infra.InfluxDB) {
	cloudKeyInformations = cloudKey
	influxDB = influx

	client := influxdb2.NewClient(
		fmt.Sprintf("http://%s:%d/", influx.Hostname, influx.Port),
		fmt.Sprintf("%s:%s", influx.Username, influx.Password))

	writeAPI := client.WriteAPIBlocking("telegraf", "telegraf")
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point immediately
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Fatalln(err)
	}
}

func Tick() {
	startTime := time.Now()
	cookie, err := login(cloudKeyInformations)

	if err != nil {
		log.Println("Unable to process tick")
		log.Println(err.Error())
	}

	health, err := GetHealth(cloudKeyInformations, cookie)
	if err != nil {
		log.Println("Unable to fetch health informations")
		log.Println(err.Error())
	}

	clients, err := GetClients(cloudKeyInformations, cookie)
	if err != nil {
		log.Println("Unable to fetch clients informations")
		log.Println(err.Error())
	}

	clientsStats, err := GetClientsStats(cloudKeyInformations, cookie)
	if err != nil {
		log.Println("Unable to fetch clients stats informations")
		log.Println(err.Error())
	}

	_, _, _ = health, clients, clientsStats
	endTime := time.Since(startTime)
	println("Tick done in: ", endTime.Milliseconds(), "ms")
}
