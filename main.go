package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	runtime.GOMAXPROCS(12)
	wg := sync.WaitGroup{}
	n := 24
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			simulateWorkload()
		}()
	}
	wg.Wait()
	log.Println("since: ", time.Since(now))
}

func simulateWorkload() {
	for i := 0; i < 1e9; i++ {
	}
}
