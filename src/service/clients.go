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
		return nil, err
	}

	var inter response.ClientsResponse
	json.NewDecoder(serverRequest.Body).Decode(&inter)

	return &inter, nil
}
