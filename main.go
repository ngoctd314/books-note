package main

import "log"

func main() {
	ch := make(chan struct{})
	log.Println(SafeClose(ch))
}

func SafeClose(ch chan struct{}) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call
			justClosed = false
		}
	}()

	close(ch)
	return true // <=> justClosed = true; return
}
