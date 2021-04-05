package main

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/service"
	"log"
)
import "github.com/davecgh/go-spew/spew"

func main() {
	infra.LoadEnv()

	cloudKey := infra.LoadCloudKey()

	cookie, err := service.Login(cloudKey)

	if err != nil {
		log.Fatalln(err)
	}

	println("COOKIE FOUND")

	//https://10.0.0.254/proxy/network/api/s/default/stat/health

	spew.Dump(service.GetHealth(cloudKey, cookie))
}
