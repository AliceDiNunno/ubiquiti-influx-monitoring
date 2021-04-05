package response

type Client struct {
	Id         string `json:"_id"`
	Mac        string `json:"mac"`
	Site       string `json:"site_id"`
	Guest      bool   `json:"is_guest"`
	FirstSeen  int    `json:"first_seen"`
	LastSeen   int    `json:"last_seen"`
	Wired      bool   `json:"is_wired"`
	Hostname   string `json:"hostname"`
	DeviceName string `json:"device_name"`
}

type ClientsResponse struct {
	Response
	Data []Client `json:"data"`
}
