package main

import "fmt"

func main() {
	chanOnwer := func() <-chan int {
		results := make(chan int)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Println(result)
		}
	}

	consumer(chanOnwer())
}
