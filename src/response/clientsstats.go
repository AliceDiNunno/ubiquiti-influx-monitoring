package response

type ClientStats struct {
	Id              string `json:"_id"`
	UserId          string `json:"user_id"`
	AssocTime       int    `json:"assoc_time"`
	LatestAssocTime int    `json:"latest_assoc_time"`
	IsWired         bool   `json:"is_wired"`
	Rssi            int    `json:"rssi"`
	Ccq             int    `json:"ccq"`
	Noise           int    `json:"noise"`
	Signal          int    `json:"signal"`
	TxPower         int    `json:"tx_power"`
	TxRetries       int    `json:"tx_retries"`
	BytesSent       int    `json:"tx_bytes"`
	BytesReceived   int    `json:"rx_bytes"`
}

type ClientsStatsResponse struct {
	Response
	Data []ClientStats `json:"data"`
}
