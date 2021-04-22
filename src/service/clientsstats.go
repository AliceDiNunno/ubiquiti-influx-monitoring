package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClientsStats(server infra.UbiquitiServer, cookie *http.Cookie) (*response.ClientsStatsResponse, error) {
	url := fmt.Sprintf("https://%s/proxy/network/api/s/%s/stat/sta", server.Hostname, server.Site)

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		serverRequest = nil
		return nil, err
	}

	var inter response.ClientsStatsResponse
	decoder := json.NewDecoder(serverRequest.Body)
	err = decoder.Decode(&inter)
	defer serverRequest.Body.Close()

	serverRequest = nil
	decoder = nil
	if err != nil {
		return nil, err
	}
	return &inter, nil
}
