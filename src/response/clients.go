package response

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/tools"
)

type Client struct {
	Id         string `json:"_id"`
	Mac        string `json:"mac"`
	Site       string `json:"site_id"`
	Guest      bool   `json:"is_guest"`
	FirstSeen  int    `json:"first_seen"`
	LastSeen   int    `json:"last_seen"`
	Wired      bool   `json:"is_wired"`
	HostName   string `json:"hostname"`
	DeviceName string `json:"device_name"`
	CustomName string `json:"name"`
}

type ClientsResponse struct {
	Response
	Data []Client `json:"data"`
}

func (c Client) GetDeviceName() string {
	names := []string{c.CustomName, c.HostName, c.DeviceName}

	for _, name := range names {
		if tools.ValidateHostName(name) {
			return name
		}
	}

	return ""
}
