package main

import (
	"fmt"
	"log"
	"time"
)

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
	if done {
		log.Println(len(a))
	}
}

func main() {
	ch := make(chan int, 4)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
