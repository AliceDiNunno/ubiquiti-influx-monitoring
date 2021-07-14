package main

import (
	"crypto/tls"
	"github.com/AliceDiNunno/ubiquiti-influx-monitoring/src/infra"
	"github.com/AliceDiNunno/ubiquiti-influx-monitoring/src/service"
	"net/http"
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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	infra.LoadEnv()
	cloudKey := infra.LoadCloudKey()
	influxConfig := infra.LoadInflux()

	instance := service.NewService(cloudKey, influxConfig)

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	waitlock.Add(1)

	catchSigInt()

	go func() {
		for {
			select {
			case <-ticker.C:
				http.DefaultTransport.(*http.Transport).CloseIdleConnections()
				instance.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	waitlock.Wait()
	println("Cleaning up...")
}
