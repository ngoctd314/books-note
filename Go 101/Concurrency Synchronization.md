# Concurrency Synchronization Techniques Provided in the sync Standard Package

In fact, channels are not the only synchronization techniques provided in Go. There are some other synchronization techniques supported by Go. For some specified circumstances, using the synchronization techniques other than channel are more efficient and readable than using channels.

The sync standard package provides several types which can be used to do synchronizations for some specialized circumstances and guarantee some specialized memory orders. For the specialized circumstances, these techniques are more efficient, and look cleaner, than the channel ways.

To avoid abnormal behaviors, it is best never to copy the values of the types in the sync standard package.

## The sync.Once Type

A *sync.Once value has a Do(f func()) method, which takes a solo parameter with type func().

For an addressable Once value o, the method call o.Do(), which is a shorthand of (&o).Do(), can be concurrently executed multiple times, in multiple goroutines. The arguments of these o.Do() calls should (but are not required to) be the same function value.

Generally, a Once value is used to ensure that a piece of code will be executed exactly once in concurrent programming.

```go
func main() {
	log.SetFlags(0)

	x := 0
	doSomething := func() {
		x++
		log.Println("Hello")
	}

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(doSomething)
			log.Println("world!")
		}()
	}

	wg.Wait()
	log.Println("x =", x) // x = 1
}
```
## The sync.Mutex and sync.RWMutex Types