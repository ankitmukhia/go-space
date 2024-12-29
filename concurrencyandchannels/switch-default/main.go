package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	// 1. This func only exits when saveAfter signals/sends data, which actually only send data after 2800 millisecond, which means it will keep running other thing till 2800 mil.
	// 2. And snapshotTicker will signals/send data, every 800 mil sec, meaning it will log "taking snapshot", every 800 mil. And keep doing every 800 mil, coz it doesn't have return.
	// 3. And at last default will keep renning ever loop, "waiting...", and in between takeSnapshot will also log, but it won't return. So it will look like, ever 800 sec data is snapshoting. And after a while it will stop(which is when saveSnapshot is true).
	for {
		select {
		case _, ok := <-snapshotTicker:
			if ok {
				takeSnapshot(logChan)
			}
		case _, ok := <-saveAfter:
			if ok {
				saveSnapshot(logChan)
				return
			}
		default: 
			waitForData(logChan)
			time.Sleep(500* time.Millisecond)
		}
	}
}
// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
