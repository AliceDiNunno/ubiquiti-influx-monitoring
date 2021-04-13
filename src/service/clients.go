package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"net/http"
)

func GetClients(server infra.UbiquitiServer, cookie *http.Cookie) (*response.ClientsResponse, error) {
	clientsEndpoint := "/proxy/network/api/s/" + server.Site + "/rest/user"

	url := "https://" + server.Hostname + clientsEndpoint

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		serverRequest = nil
		return nil, err
	}

	var inter response.ClientsResponse
	decoder := json.NewDecoder(serverRequest.Body)
	err = decoder.Decode(&inter)

	decoder = nil
	serverRequest = nil
	if err != nil {
		return nil, err
	}
	return &inter, nil
}
