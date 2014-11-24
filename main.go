package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/cloudfoundry/noaa"
	"github.com/cloudfoundry/noaa/events"
	"github.com/pivotal-cf-experimental/veritas/say"
)

func main() {
	if len(os.Args) != 2 {
		PrintUsageAndExit()
	}

	logGuid := os.Args[1]
	address := os.Getenv("DOPPLER")
	if address == "" {
		address = "wss://doppler.ketchup.cf-app.com:4443"
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	logConsumer := noaa.NewConsumer(address, tlsConfig, nil)
	outputChan := make(chan *events.LogMessage, 0)
	errorChan := make(chan error, 0)
	stopChan := make(chan struct{}, 0)
	go logConsumer.TailingLogs(logGuid, "", outputChan, errorChan, stopChan)

	say.Println(0, "Fetching logs for log-guid: %s", say.Green(logGuid))

	for {
		select {
		case message := <-outputChan:
			t := time.Unix(0, message.GetTimestamp())
			say.Println(0, "%s [%s|%s] %s", say.Green(t.Format("02 Jan 15:04")), say.Cyan("%s", message.GetSourceInstance()), say.Cyan("%s", message.GetSourceType()), string(message.GetMessage()))
		case err := <-errorChan:
			say.Println(0, say.Red("Error while streaming:\n%s", err.Error()))
			return
		}
	}

}

func PrintUsageAndExit() {
	fmt.Println(`Usage:
picard LOG-GUID

Set the loggregator address with the DOPPLER environment:

    export DOPPLER=wss://doppler.ketchup.cf-app.com:4443

Defaults to ketchup without the environment variable.
The address for a local Diego Edge box can be set via: 

    export DOPPLER=ws://doppler.192.168.11.11.xip.io:443 //CHECK THIS!!
`)
	os.Exit(1)
}
