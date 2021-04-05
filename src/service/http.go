package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

func buildClient() *http.Client {
	client := &http.Client{}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return client
}

func buildRequest(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("Accept", "*/*")

	return request, nil
}

func httpGET(url string, cookie *http.Cookie) (*http.Response, error) {
	client := buildClient()
	request, err := buildRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if cookie != nil {
		request.AddCookie(cookie)
	}

	return client.Do(request)

}

func httpPOST(url string, body interface{}, cookie *http.Cookie) (*http.Response, error) {
	bodyJson, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	client := buildClient()
	request, err := buildRequest("POST", url, bytes.NewBuffer(bodyJson))

	if err != nil {
		return nil, err
	}

	if cookie != nil {
		request.AddCookie(cookie)
	}

	return client.Do(request)
}
