package main

import (
	"fmt"
	"time"
)

type T struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

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
