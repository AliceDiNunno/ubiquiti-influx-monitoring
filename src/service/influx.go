package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

const ( //Telegraf compatible input names
	InputBytesSent           = "tx_bytes"
	InputBytesReceived       = "rx_bytes"
	InputTransmissionRetried = "tx_retried"
	InputTransmissionPower   = "tx_power"
	InputIsWired             = "is_wired"
	InputPing                = "ping"

	InputWlanCCQ = "wlan_ccq"
	InputNoise   = "wlan_noise"
	InputRssi    = "wlan_rssi"
	InputSignal  = "wlan_signal"

	InputSwitchCount              = "count_switch"
	InputApCount                  = "count_ap"
	InputAdoptedDevicesCount      = "count_adopted_devices"
	InputDisconnectedDevicesCount = "count_disconnected_devices"
	InputGuestCount               = "count_guest"
	InputPendingDevicesCount      = "count_pending_devices"
	InputIoTCount                 = "count_iot"
	InputUserCount                = "count_user"
)

func sendHealthMetrics(influx influxdb2.Client, metrics []response.Health) {
	for _, _ = range metrics {
		//TO DO

		/*println(healthMetric.Subsystem)
		println(InputBytesSent, healthMetric.BytesSent)
		println(InputBytesReceived, healthMetric.BytesReceived)
		println(InputSwitchCount, healthMetric.SwitchCount)
		println(InputApCount, healthMetric.AccessPointCount)
		println(InputAdoptedDevicesCount, healthMetric.AdoptedDevicesCount)
		println(InputDisconnectedDevicesCount, healthMetric.DisconnectedDevicesCount)
		println(InputGuestCount, healthMetric.GuestCount)
		println(InputPendingDevicesCount, healthMetric.PendingDevicesCount)
		println(InputIoTCount, healthMetric.IoTCount)
		println(InputUserCount, healthMetric.UserCount)
		println(InputPing, healthMetric.SpeedTestPing)

		println("=====")*/

	}
}

func newPoint(client response.Client, tag string) *write.Point {
	p := influxdb2.NewPointWithMeasurement(tag)

	p = p.AddTag("host", client.GetDeviceName())
	p = p.AddTag("id", client.Id)

	p.SetTime(time.Now())

	return p
}

func newNetPoint(client response.Client) *write.Point {
	return newPoint(client, "net")
}

func newWlanPoint(client response.Client) *write.Point {
	return newPoint(client, "wlan")
}

func sendDeviceMetrics(influx influxdb2.Client, metrics map[response.Client]response.ClientStats) {
	points := []*write.Point{}
	writeAPI := influx.WriteAPIBlocking("telegraf", "telegraf")

	for client, stat := range metrics {
		netPoint := newNetPoint(client)

		if !stat.IsWired {
			wlanPoint := newWlanPoint(client)
			wlanPoint = wlanPoint.AddField(InputWlanCCQ, stat.Ccq)
			wlanPoint = wlanPoint.AddField(InputNoise, stat.Noise)
			wlanPoint = wlanPoint.AddField(InputRssi, stat.Rssi)
			wlanPoint = wlanPoint.AddField(InputSignal, stat.Signal)
			wlanPoint = wlanPoint.AddField(InputTransmissionPower, stat.TxPower)
			points = append(points, wlanPoint)
		}

		netPoint = netPoint.AddField(InputBytesReceived, stat.BytesReceived)
		netPoint = netPoint.AddField(InputBytesSent, stat.BytesSent)
		netPoint = netPoint.AddField(InputTransmissionRetried, stat.TxRetries)
		points = append(points, netPoint)
	}

	// write all the points
	err := writeAPI.WritePoint(context.Background(), points...)
	if err != nil {
		log.Fatalln(err)
	}
}
