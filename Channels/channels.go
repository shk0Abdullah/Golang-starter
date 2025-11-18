// channels:
// channels They were acting like a pipe You sends data from one side of the pipe and receive it from the other side of the side
// When Your multiple goroutines running and you want one goroutine to give specific data to the other one As They were doing tasks
// concurrently You can't do all by using simple methods So here Channels concepts work
package main

import (
	"fmt"
	"time"
)

// This will throw Deadlock in program error
// Deadlock in OS:
// Deadlock is when multiple tasks run concurrently And One task is getting hold of one resource and the other task is waiting for the
// release of the resource and the first one Why its not releasing the resource cause the other one holding another dependent resource
// and that task is waiting to get the release of the resource This is Deadlock concept They were in the infinite deadlock

func processNum(numchan chan int) {
	msgreceived := <-numchan
	fmt.Println("Receiving this number:", msgreceived)

}

func main() {
	// messageChnl := make(chan string)
	// The below ones are blocking
	// its blocking until the second side is not ready to receive the data
	// messageChnl <- "ping" // sending data to channel
	// msg := <-messageChnl  // receiving data
	// fmt.Println(msg)
	numChnl := make(chan int)
	// The below is not blocking it will run and exit immediately use time or waitgroups
	go processNum(numChnl)
	// Must do this after making a groutine
	numChnl <- 28

	time.Sleep(time.Second * 2)

}
