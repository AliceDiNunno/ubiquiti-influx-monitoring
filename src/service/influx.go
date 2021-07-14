package service

import (
	"context"
	"github.com/AliceDiNunno/gobiquiti"
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

func (i *Instance) sendHealthMetrics(metrics []gobiquiti.Health) {
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
	metrics = nil
}

func newPoint(client gobiquiti.Client, tag string) *write.Point {
	p := influxdb2.NewPointWithMeasurement(tag)

	p = p.AddTag("host", client.GetDeviceName())
	p = p.AddTag("id", client.Id)

	p.SetTime(time.Now())

	return p
}

func newNetPoint(client gobiquiti.Client) *write.Point {
	return newPoint(client, "net")
}

func newWlanPoint(client gobiquiti.Client) *write.Point {
	return newPoint(client, "wlan")
}

func (i *Instance) sendDeviceMetrics(metrics map[gobiquiti.Client]gobiquiti.ClientStats) {
	var points []*write.Point
	writeAPI := i.influxClient.WriteAPIBlocking("telegraf", "telegraf")

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
			wlanPoint = nil
		}

		netPoint = netPoint.AddField(InputBytesReceived, stat.BytesReceived)
		netPoint = netPoint.AddField(InputBytesSent, stat.BytesSent)
		netPoint = netPoint.AddField(InputTransmissionRetried, stat.TxRetries)
		points = append(points, netPoint)
		netPoint = nil
	}

	// write all the points
	err := writeAPI.WritePoint(context.Background(), points...)
	if err != nil {
		log.Fatalln(err)
	}

	for _, pointToFree := range points {
		_ = pointToFree
		pointToFree = nil
	}

	points = nil
}
