package service

import (
	"github.com/AliceDiNunno/gobiquiti"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Instance struct {
	CloudKey     gobiquiti.CloudKeyInstance
	influxClient influxdb2.Client
}
