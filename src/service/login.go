package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/requests"
	"errors"
	"net/http"
	"time"
)

var cookie *http.Cookie = nil
var cookieLastFetch time.Time

func cookieRefreshRequired() bool {
	if cookie == nil {
		return true
	}

	return time.Since(cookieLastFetch).Minutes() > 50 //Ubiquiti's JWT elapses an hour after generation
}

func login(server infra.UbiquitiServer) (*http.Cookie, error) {
	if !cookieRefreshRequired() {
		return cookie, nil
	}

	cookie = nil

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
		cookieLastFetch = time.Now()
		return request.Cookies()[0], nil
	}

	return nil, errors.New("Server did not respond with a valid token")
}
