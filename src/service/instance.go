package service

import (
	"fmt"
	"github.com/AliceDiNunno/gobiquiti"
	"github.com/AliceDiNunno/ubiquiti-influx-monitoring/src/infra"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Instance struct {
	CloudKey     gobiquiti.CloudKeyInstance
	influxClient influxdb2.Client

	InfluxOrg    string
	InfluxBucket string
}

func NewInstance(cloudKey gobiquiti.Config, influx infra.InfluxDB) *Instance {
	cloudKeyInstance := gobiquiti.CloudKeyInstance{Config: cloudKey}

	protocol := "http"

	if influx.IsSecure {
		protocol = "https"
	}

	influxClient := influxdb2.NewClient(
		fmt.Sprintf("%s://%s:%d/", protocol, influx.Hostname, influx.Port),
		influx.Token)

	return &Instance{
		CloudKey:     cloudKeyInstance,
		influxClient: influxClient,

		InfluxOrg:    influx.Organization,
		InfluxBucket: influx.Bucket,
	}
}
