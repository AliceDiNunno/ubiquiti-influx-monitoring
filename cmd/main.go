package main

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/service"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var waitlock = &sync.WaitGroup{}

func catchSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == syscall.SIGINT {
				waitlock.Done()
			}
		}
	}()
}

func main() {
	infra.LoadEnv()
	cloudKey := infra.LoadCloudKey()
	influxConfig := infra.LoadInflux()

	service.LoadService(cloudKey, influxConfig)

	/*

		cloudKey := infra.LoadCloudKey()

		cookie, err := service.Login(cloudKey)

		if err != nil {
			log.Fatalln(err)
		}

		println("COOKIE FOUND")

		//https://10.0.0.254/proxy/network/api/s/default/stat/health

		spew.Dump(service.GetClientsStats(cloudKey, cookie))*/

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	waitlock.Add(1)

	catchSigInt()

	go func() {
		for {
			select {
			case <-ticker.C:
				service.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	waitlock.Wait()
	println("Cleaning up...")
}
