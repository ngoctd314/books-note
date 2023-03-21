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
	now := time.Now()
	notify := make(chan struct{})

	go func() {
		<-notify
		fmt.Println("since: ", time.Since(now))
	}()
	go func() {
		<-notify
		fmt.Println("since: ", time.Since(now))
	}()

	time.Sleep(time.Second)
	close(notify)
	time.Sleep(time.Second)
}
