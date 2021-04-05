package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"log"
	"net/http"
)

func panicError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetHealth(server infra.UbiquitiServer, cookie *http.Cookie) (*response.HealthResponse, error) {
	healthEndpoint := "/proxy/network/api/s/default/stat/health"

	url := "https://" + server.Hostname + healthEndpoint

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		return nil, err
	}

	var inter response.HealthResponse
	json.NewDecoder(serverRequest.Body).Decode(&inter)

	return &inter, nil
}
