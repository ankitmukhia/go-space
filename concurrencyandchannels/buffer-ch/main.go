package main

// https://www.boot.dev/lessons/d2a10614-3142-4d3e-906e-5a817aa920b3
// batching: refers to the process of gruping multipe individual items

// As we can see we are using channel here without go keyword(go routine). But this problem statment is uses buffer channel, means we can send to the buffer, even if there is no reciver. So we should be able to do it in the same thread. Only thing is we need enough batch size(which is buffer size). 

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	for _, em := range emails {
		ch <- em		
	}
	return ch 
}
