package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"net/http"
)

func GetHealth(server infra.UbiquitiServer, cookie *http.Cookie) (*response.HealthResponse, error) {
	healthEndpoint := "/proxy/network/api/s/" + server.Site + "/stat/health"

	url := "https://" + server.Hostname + healthEndpoint

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		return nil, err
	}

	var inter response.HealthResponse
	decoder := json.NewDecoder(serverRequest.Body)
	err = decoder.Decode(&inter)
	defer serverRequest.Body.Close()

	decoder = nil
	serverRequest = nil
	if err != nil {
		return nil, err
	}
	return &inter, nil
}
