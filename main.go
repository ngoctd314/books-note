package main

import (
	"log"
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
