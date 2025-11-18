package main

import "fmt"

func sum(result chan int, no1 int, no2 int) {
	result <- no1 + no2
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing")
}

func main() {

	done := make(chan bool)
	go task(done)
	// We can do the same using channels instead of Waitgroup
	// When you have multiple channels use Waitgroup
	val := <-done //blocking
	fmt.Println(val)
	// result := make(chan int)

	// go sum(result, 4, 5)
	// // sending and receiving are blocking when we use goroutines excluding buffers
	// receive := <-result
	// fmt.Println(receive)

}
