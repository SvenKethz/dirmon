package main

import (
	"fmt"
	"time"
)

var fahrrunter = make(chan bool, 1)

func runMonitoring(directory string, idleTime int) {
	for {
		select {
		case <-fahrrunter:
			return
		default:
			monitor(directory)
			catchSignalsWhileIdle(idleTime)
		}
	}
}

func catchSignalsWhileIdle(idleTime int) {
	for i := 1; i <= idleTime; i++ {
		select {
		case <-stopchannel:
			fmt.Println("\nrecieved stop signal, will shut down")
			fahrrunter <- true
			fmt.Println("waiting for processes to finish...")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
