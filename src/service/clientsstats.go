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
		return nil, err
	}

	var inter response.ClientsStatsResponse
	json.NewDecoder(serverRequest.Body).Decode(&inter)
	serverRequest = nil

	return &inter, nil
}
