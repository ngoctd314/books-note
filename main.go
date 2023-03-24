package main

import (
	"fmt"
	"runtime"
)

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()

		return completed
	}
	doWork(nil) // block for ever (receiving value from nil channel)
	// Perhaps more work is done here
	fmt.Println("Done.")
	fmt.Println(runtime.NumGoroutine())
}

func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}

func sum8(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 8 {
		total += s[i]
	}
	return total
}
