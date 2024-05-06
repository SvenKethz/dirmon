package main

import (
	"time"
)

var (
	Fahrrunter = make(chan bool, 1)
	NewFiles   = make(chan []string)
)

func monitor(directory string, idleTime int) {
	for i := 1; i <= idleTime; i++ {
		select {
		case <-Fahrrunter:
			return
		default:
			checkDir(directory)
			time.Sleep(1 * time.Second)
		}
	}
}

func checkDir(directory string) {
	filelist := []string{"Hallo", "Welt", directory}

	NewFiles <- filelist
}
