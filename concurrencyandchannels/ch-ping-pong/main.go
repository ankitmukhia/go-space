package main

import (
	"fmt"
	"time"
)
// Problem solution explanation: The self-calling function will work when we remove it as a goroutine. 
// What happens when we make it a goroutine is that it runs concurrently with the rest of the program, 
// meaning the main program does not wait for the goroutine to finish because it doesn't communicate 
// via channels, which block the process until the sender and receiver are synchronized with each other.

func pingPong(numPings int) {
	pings := make(chan struct{}) // to send pings msg.
	pongs := make(chan struct{}) // to send pongs msg.
	go pinger(pings, pongs, numPings) // sends 4 "ping" msg.
	go ponger(pings, pongs) // responds to "ping" msg with "pong".
	go func (){
		// when the pongs chan is closed, the loop ends, and prints "pongs done"
		i := 0 
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}

// don't touch below this line

func pinger(pings, pongs chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond // waits for given duration the double each time
	// if "numPings" is 4, then after 400ms, the loop exit, and pings chan will close
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	// purpose of this func is to respond to every "ping" with "pong"
	// listen to pings chan, for range, such that, it will respond to each ping.
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	// when ping chan is closed and loop ends and print "pings done"
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
