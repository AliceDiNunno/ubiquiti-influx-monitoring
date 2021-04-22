package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHealth(server infra.UbiquitiServer, cookie *http.Cookie) (*response.HealthResponse, error) {
	url := fmt.Sprintf("https://%s/proxy/network/api/s/%s/stat/health", server.Hostname, server.Site)

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
