package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)
	// consume
	go func() {
		for range ch {
			time.Sleep(time.Second)
			fmt.Println("consume")
			<-ch
		}
	}()

	for i := 0; i < 100; i++ {
		fmt.Println("send")
		ch <- struct{}{}
	}

	time.Sleep(time.Minute)
}
