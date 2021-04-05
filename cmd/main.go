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
