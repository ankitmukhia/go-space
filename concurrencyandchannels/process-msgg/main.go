package main

import "time"

func processMessages(messges []string) []string {
	ch := make(chan string)
	processedMessages := []string{}
	// this was my first sol, but it has issue of deadlock. Means code immediately ties to receive from the chan before msg was even sent
	// processedMessages = append(processedMessages, <-ch) 

	go func() {
		for _, mesg := range messges {
			ch <- process(mesg)
		}
		close(ch)
	}()
	// collect all the processed messages from the chan
	for processedMsg := range ch {
	  processedMessages = append(processedMessages, processedMsg) 
	}
	return processedMessages
}

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
