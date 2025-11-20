// There is no mean of Multi-threading if at the end you locking the resources for a single request to complete the modifictaion
// Best practice is you should lock the only modification logic and leave the rest out of the lock to get benefit from
// concurrency
package main

import (
	"fmt"
	"sync"
)

// Mutex: Mutual Exclusion
// When we use Multi threading to get rid of Race condition we use Mutex

type post struct {
	// Because of goroutine multiple processes try to modify when clash they overwrite each other's instructions That's why you would see a no less than 100
	// sometimes
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	// locking so only one request modifies it at a time
	p.mu.Lock()
	p.views += 1

}

func race() {
	var wg sync.WaitGroup
	Mypost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Mypost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Views:", Mypost.views)
}

func main() {
	// This func is creating race condition when goroutines tries to concurrently use the same resource struct's views
	// They clash and end up overwriting each others processes
	race()
	// Solution: Mutex
	// we lock that resource to get modified once at a time
}
