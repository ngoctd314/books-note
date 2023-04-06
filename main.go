package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int)
	// stopCh is an additional signal channel
	// Its sender is the moderator goroutine, and its receivers are all senders and receiver
	stopCh := make(chan struct{})
	// the channel toStop is used to notify the moderator to close the additional signal channel (stopCh)
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// Here, try-send operator is to notify the moderator to close the additional signal channel
					// we can't close(stopCh) here because concurrency and close closed channel problem
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to try to exit the sender goroutine as early as possible. Try-receiver and try-send select blocks are specially optimized by the standard go compiler
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first brand in this select block might be still not selected for some loops if the the send to dataCh is also non-blocking. If this is unacceptable,
				// then the above try-receive operation is essential
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				select {
				// Same as the sender goroutine, the try-receive operation here is to try to exit the receiver goroutine as early as possible
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and forever in theory) if the receive
				// from dataCh is also non-blocking. If
				// this is not acceptable, then the above
				// try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
					}
					log.Println(value)
				}
			}

		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
