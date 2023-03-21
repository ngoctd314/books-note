package main

import (
	"fmt"
	"time"
)

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
