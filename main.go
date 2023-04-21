package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"sync"
)

func main() {
	log.SetFlags(0)

	x := 0
	fn := func() {
		x++
	}
	oc := sync.Once{}
	for i := 0; i < 1e6; i++ {
		go oc.Do(fn)
	}
	fmt.Println(x)
}
