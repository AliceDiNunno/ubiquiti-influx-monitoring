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
	json.NewDecoder(serverRequest.Body).Decode(&inter)
	serverRequest = nil

	return &inter, nil
}
