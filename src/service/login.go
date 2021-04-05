package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/requests"
	"errors"
	"net/http"
)

func Login(server infra.UbiquitiServer) (*http.Cookie, error) {
	loginEndpoint := "/api/auth/login"

	url := "https://" + server.Hostname + loginEndpoint

	login := requests.UbiquitiLoginRequest{
		Username: server.Username,
		Password: server.Password,
	}

	request, err := httpPOST(url, login, nil)

	if err != nil {
		return nil, err
	}

	if len(request.Cookies()) > 0 {
		return request.Cookies()[0], nil
	}

	return nil, errors.New("Server did not respond with a valid token")
}
