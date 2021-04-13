package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"net/http"
)

func GetClientsStats(server infra.UbiquitiServer, cookie *http.Cookie) (*response.ClientsStatsResponse, error) {
	clientsEndpoint := "/proxy/network/api/s/" + server.Site + "/stat/sta"

	url := "https://" + server.Hostname + clientsEndpoint

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		serverRequest = nil
		return nil, err
	}

	var inter response.ClientsStatsResponse
	decoder := json.NewDecoder(serverRequest.Body)
	err = decoder.Decode(&inter)

	serverRequest = nil
	decoder = nil
	if err != nil {
		return nil, err
	}
	return &inter, nil
}
