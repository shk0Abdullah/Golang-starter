package main

import (
	"fmt"
	"time"
)

// The most hats off feature of golang for which I was learning Golang
// Goroutines:
// This is used when you want to implement multithreading by doing concurrent tasks

func task(id int) {
	fmt.Println("doinng tasks:", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// write go to perform concurrent tasks depending upon the cors of your PC
		// go task(i)
		// anonymous functiion with closure
		// its a good practice to receive the required use variables within the closures
		go func() { fmt.Println(i) }()
		// With receiving vars both execute
		go func(i int) { fmt.Println(i) }(i)

	}
	time.Sleep(time.Second * 2)
}

// Very fast as compared to blocking
