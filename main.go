package main

import (
	"time"
)

var (
	Fahrrunter = make(chan bool, 1)
	NewFiles   = make(chan []string)
)

func Monitor(directory string, idleTime int) {
	for {
		select {
		case <-Fahrrunter:
			return
		default:
			CheckDir(directory)
			time.Sleep(time.Duration(idleTime) * time.Second)
		}
	}
}

func CheckDir(directory string) {
	var filelist []string
	existing_files := 0
	if len(filelist) != existing_files {
		existing_files = len(filelist)
		NewFiles <- filelist
	} else {
		filelist = []string{"Hallo", "Welt", directory}
	}
}
