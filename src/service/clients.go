package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetClients(server infra.UbiquitiServer, cookie *http.Cookie) (*response.ClientsResponse, error) {
	url := fmt.Sprintf("https://%s/proxy/network/api/s/%s/rest/user", server.Hostname, server.Site)

	serverRequest, err := httpGET(url, cookie)

	if err != nil {
		serverRequest = nil
		return nil, err
	}

	var inter response.ClientsResponse
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
