package main 

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// here order should be 1st someone send, then other receive.
	// but right now order is swaped, 1st receive, 2nd send. which is wrong //
	// go keyword make this anonymous func run concurently. 
	go func() {
		time.Sleep(time.Millisecond * 250)	
		fmt.Printf("Email received: '%s'\n", message)
	}()	
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string){
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("=======================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}
