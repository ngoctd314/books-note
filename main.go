package main

import (
	"fmt"
	"sync"
)

func main() {
}

func incorrectSync() {
	var (
		x, y int
		wg   sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 1
		fmt.Print("y: ", y, " ")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		y = 1
		fmt.Print("x: ", x, " ")
	}()

	wg.Wait()
}
