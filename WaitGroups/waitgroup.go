package main

import (
	"fmt"
	"sync"
)

// The most hats off feature of golang for which I was learning Golang
// Goroutines:
// This is used when you want to implement multithreading by doing concurrent tasks

func task(id int, w *sync.WaitGroup) {
	// defer runs at the end of the execution function
	defer w.Done()
	fmt.Println("doinng tasks:", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) { task(i, &wg) }(i)

	}
	// time.Sleep(time.Second * 2)
	wg.Wait()
}

// Very fast as compared to blocking
