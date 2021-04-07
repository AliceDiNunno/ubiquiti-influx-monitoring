package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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
	for _, healthMetric := range metrics {
		println(healthMetric.Subsystem)
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

		println("=====")

	}
}

func sendDeviceMetrics(influx influxdb2.Client, metrics map[response.Client]response.ClientStats) {
	for client, stat := range metrics {
		_ = stat

		writeAPI := influx.WriteAPIBlocking("telegraf", "telegraf")

		p := influxdb2.NewPointWithMeasurement("system")

		p = p.AddTag("host", client.GetDeviceName())
		p = p.AddTag("id", client.Id)

		p = p.AddField(InputIsWired, stat.IsWired)
		if !stat.IsWired {
			p = p.AddField(InputWlanCCQ, stat.Ccq)
			p = p.AddField(InputNoise, stat.Noise)
			p = p.AddField(InputRssi, stat.Rssi)
			p = p.AddField(InputSignal, stat.Signal)
			p = p.AddField(InputTransmissionPower, stat.TxPower)
		}
		p = p.AddField(InputBytesReceived, stat.BytesReceived)
		p = p.AddField(InputBytesSent, stat.BytesSent)
		p = p.AddField(InputTransmissionRetried, stat.TxRetries)

		p.SetTime(time.Now())

		// write point immediately
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
