// Buffered channels we can send more than one value
// Limited amount of date transfer between goroutines

// Buffered
// Think a channel like a stack having capacity of 1 by default if you put the value getting the value out will release the
// channel which is blocking

package main

import (
	"fmt"
)

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		// time.Sleep(time.Second)
	}
}

func emailCaller() {
	emailChan := make(chan string, 100)
	done := make(chan bool)

	for i := 0; i < 100; i++ {
		email := fmt.Sprintf("%d@example.com", i)
		emailChan <- email
	}
	go emailSender(emailChan, done)
	// Important to close the Channels System
	close(emailChan)
	<-done
}

func starter() {
	emailChan := make(chan string, 100)

	emailChan <- "1@example.com"
	fmt.Println(<-emailChan)
	// Nothing in the channel error
	// fmt.Println(<-emailChan)
	emailChan <- "2@example.com"
	emailChan <- "2@example.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
}

func multipleChan() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 10 }()
	go func() { chan2 <- "pong" }()

	for i := 0; i < 3; i++ {
		select {
		case Chan1Val := <-chan1:
			fmt.Println("Chan1's value", Chan1Val)
		case Chan2Val := <-chan2:
			fmt.Println("Chan2's value", Chan2Val)

		}
	}
	close(chan1)
	close(chan2)
}

func main() {
	// starter()
	// emailCaller()
	// at a time listening to multiple channels
	multipleChan()

}

// Receive only Channels
// emailSender(emailChan <-chan)
// Send only Channels
// emailSender(emailChan chan<-)
