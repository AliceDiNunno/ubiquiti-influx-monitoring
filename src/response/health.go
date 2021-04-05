package response

type Health struct {
	Subsystem                string `json:"subsystem"`
	UserCount                int    `json:"num_user"`
	GuestCount               int    `json:"num_guest"`
	IoTCount                 int    `json:"num_iot"`
	BytesSent                int    `json:"tx_bytes-r"`
	BytesReceived            int    `json:"rx_bytes-r"`
	Status                   string `json:"status"`
	AccessPointCount         int    `json:"num_ap"`
	SwitchCount              int    `json:"num_sw"`
	AdoptedDevicesCount      int    `json:"num_adopted"`
	DisconnectedDevicesCount int    `json:"num_disconnected"`
	PendingDevicesCount      int    `json:"num_pending"`
	Uptime                   int    `json:"uptime"`
	LanIp                    string `json:"lan_ip"`
	WanIp                    string `json:"wan_ip"`
	LastSpeedTest            int    `json:"speedtest_lastrun"`
	SpeedTestPing            int    `json:"speedtest_ping"`
}

type HealthResponse struct {
	Response
	Data []Health `json:"data"`
}
