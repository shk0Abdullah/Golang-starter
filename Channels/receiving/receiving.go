package main

import "fmt"

func sum(result chan int, no1 int, no2 int) {
	result <- no1 + no2
}

func main() {

	result := make(chan int)

	go sum(result, 4, 5)
	// sending and receiving are blocking when we use goroutines excluding buffers
	receive := <-result
	fmt.Println(receive)

}
