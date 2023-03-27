# Concurrency in Go

## Chapter 1. An Introduction to Concurrency

Concurrency is an interesting word because it means different things to different people in our field. In addition to "concurrency", you may have heard the words, "asynchronous", "parallel" or "threaded".

**Moore's Law, Web Scale, and the Mess We're In**

For problems that are embarrassingly parallel, it is recommended that you write your application so that it can scale horizontally. This means that you can take instances of  your program, run it on more CPUs, or machines, and this will cause the runtime of the system to improve.

**Why Is Concurrency Hard?**

**Race Conditions**

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

Most of the time, this shows up in what's called a data race, where one concurrent operation attempts to read a variable while at some undertermined time another concurrent operation is attempting to write to the same variable.

```go
var data int
go func() {
    data++ // line 3
}()
if data == 0 { // line 5
    fmt.Printf("the value is %v.\n", data) // line6
}
```

There are three possible outcomes to running this code:
- Nothing is printed. Stack:  line 3 -> line 5
- Print 0. Stack: line 5 -> line 6 -> line3 (or exist)
- Print 1. Stack: line 5 -> line 3 -> line 6

**Atomicity**

When something is considered atomic, or to have the property of atomicity, this means that within the context that it is operating, it is indivisible, or uninterruptible. In other words, the atomicity of an operation can change depending on the currently defined scope.

Example
```go
i++
```

It may look atomic, but a bried analysis reveals several operations:
- Retrieve the value of i
- Increment the value of i
- Store the value of i

While each of these operations alone is atomic, the combination of the three may not be, depending on your context. This reveals an interesting property of atomic operations: combining them does not necessarily produce a larger atomic operation.

Atomicity is important because if something is atomic, implicitly it is safe within concurrent contexts. This allows us to compose logically correct programs.

Most statements are not atomic, let alone functions, methods and programs. If atomicity is the key to composing logically correct programs, and most statements aren't atomic, how de we reconcile these two statements?

**Memory Access Synchronization**

Let's say we have a data race: two concurrent processes are attempting to access the same area of memory, and the way they accessing the memory is not atomic.

```go
var data int
go func() {data++}()
if data == 0 {
    fmt.Println("the value is 0.")
} else {
    fmt.Println("the value is %v.\n", data)
}
```
In fact, there's a name for a section of your program that needs exclusive access to a shared resource. This is called a critical section. In this example, we have three critical sections:

- Our goroutine, which is incrementing the data variables.
- Our if statement, which checks whether the value of data is 0.
- Our fmt.Printf statement, which retrieves the value of data for output.

One way to solve this problem is to synchronize access to the memory between your critical section.

```go
var (
    data int
    m sync.Mutex
)

go func() {
    m.Lock()
    defer m.Unlock()
    data++
}()


m.Lock()
if data == 0 {
    fmt.Println("the value is 0.")
} else {
    fmt.Println("the value is %v.\n", data)
}
m.Unlock()

```

We have solved our data race, we haven't actually solved our race condition! The order of operations in this program is still nondeterministic; we've just narrowed the scope of the nondeterminism a bit.

Solve data race is very easy, you can solve some problems by synchronizing access to the memory, but as we just say, it doesn't automatically solve data races or logical correctness. Further, it can also create maintenance and performance problems.

Synchronizing access to the memory in this manner also has performance ramifactions. This brings up two questions:
- Are my critical sections entered and exited repeatedly?
- What size should my critical sections be?

**Deadlocks, Livelocks and Starvation**

**Deadlock**

A deadlocked program is one i which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.

```go
func main() {
	m1 := sync.Mutex{}
	m2 := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		m1.Lock()
		defer m1.Unlock()

		time.Sleep(time.Second)

		m2.Lock()
		defer m2.Unlock()
	}()

	go func() {
		defer wg.Done()
		m2.Lock()
		defer m2.Unlock()

		time.Sleep(time.Second)

		m1.Lock()
		defer m1.Unlock()
	}()

	wg.Wait()
}
```

Coffman Condition
Mutual Exclusion: A concurrent process holds exclusive rights to a resource at any one time.
Wait For Condition: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
No Preemption: A resources held by a concurrent process can only be released by that process.
Circular Wait: A concurrent process (P1) must be waiting on a chain of other concurrent process (P2), which are in turn waiting on it (P1).

The laws allow us to prevent deadlocks too. If we ensure that at least one of these conditions is not true, we can prevent deadlocks from occurring.

**Live lock**

Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

**Starvation**

Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.

## Chapter 2. Modeling Your Code: Communicating Sequential Processes.

**What is CSP?**

## Chapter 3: Go's Concurrency Building Blocks

What's happening behind the scenes here: how do goroutines actually work? Are they OS threads? Green threads? How many can we create?

Goroutines are unique to Go, they're not OS threads, and they're not exactly green threads - threads that are managed by a language's runtime - they're a higher level of abstraction known as coroutines. Coroutines are simply concurrent subroutines (function, closure, or methods in Go).

### The GOMAXPROCS Lever

In the runtime package, there is a function GOMAXPROCS. The name is misleading: people often think this function relates to the number of logical processors on the host machine but really this function controls the number of OS threads that will host so-called "work queues".


## Chapter 4: Concurrency Patterns in Go

### Confinement

When working with concurrent code, there are a few different options for safe operation. We've gone over two of them:
- Synchronization primitives for sharing memory (sync.Mutex)
- Synchronization via communicating (channels)

However, there are a couple of other options that are implicitly safe within multiple concurrent processes:
- Immutable data
- Data protected by confinement

In some sense, immutable data is ideal because it is implicitly concurrent-safe. Each concurrent process may operate on the same data, but it may not modify it. If it wants to create new data, it must create a new copy of the data with the desired modification.

Confinement can also allow for a lighter cognitive load on the developer and smaller critical sections. The techniques to confine concurrent values are a bit more involved than simply passing copies of values. Confinement  is the simple yet powerful idea of ensuring information is only ever available from one concurrent process. When this is achieved, a concurrent program is implicitly safe and no synchronization is needed. There are two kinds of confinement possible: ad hoc and lexical.

```go
func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
```

We can see that the data slice of integers is available from both the loopData function and the loop over the handleData channel; however, by convention we're only accessing it from the loopData function. But as the code is touched by many people, and deadlines loom, mistakes might be made, and the confinement might break down and cause issues.

Lexical confinement involves using lexical scope to expose only the correct data and concurrency primitives for multiple concurrent processes to use. It makes it impossible to do the wrong thing.

```go
chanOwner := func() <-chan int {
	// we instantiate the channel within the lexical scope of the channel within the lexical scope of the chanOwner function.
	// This limits the scope of the write aspect of the results channel to the closure defined below it. In other words, it confines
	// the write aspect of this channel to prevent other goroutines from writing it.
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 0; i < 5; i++ {
			results <- i
		}
	}()

	return results
}
// Here we receive a read-only copy of an int channel.
// By declaring that we only usage we require is read access, we confine usage of the channel within the consume function to only reads.
consumer := func(results <-chan int) {
	for result := range results {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving")
}

// Here we receive the read aspect of the channel and we're able to pass it into the consumer, which can do nothing but read from it
// Once again this confines the main goroutine to a read-only view of the channel
results := chanOwner()
consumer(results)
```

### Preventing Goroutine Leaks

Let's start with a simple example of a goroutine leak:

```go
doWork := func(strings <- chan string) <- chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited")
		defer close(completed)
		for s := range strings {
			// Do something interesting
			fmt.Println(s)
		}
	}()

	return completed
}
doWork(nil) // block for ever (receiving value from nil channel)
// Perhaps more work is done here
fmt.Println("Done.")
fmt.Println(runtime.NumGoroutine())
```
Here we see that the main goroutine passes a nil channel into doWork. Therefore, the strings channel will never actually gets any strings written onto it, and the goroutine containing doWork will remain in memory for the lifetime of this process.

In this example, the lifetime of the process is very short, but in a real program, goroutines could easily be started at the beginning of a long-lived program. In the worst case, the main goroutine could continue to spin up goroutines throughout its life causing creep in memory utilization.

The way to successfully, mitigate this is to establish a signal between the parent goroutine and its children that allows the parent to signal cancellation to its children. By convention, this signal is usually a read-only channel named done. The parent goroutine passes this channel to the child goroutine and then closes the channel when it wants to cancel the child goroutine.

```go
func main() {
	doWork := func(cancel <-chan bool, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					completed <- s
				case s := <-strings:
					// Do something interesting
					completed <- s
				case s := <-strings:
					// Do something interesting
					completed <- s
				case s := <-strings:
					// Do something interesting
					completed <- s
				case s := <-strings:
					// Do something interesting
					completed <- s
				case <-cancel:
					completed <- "cancel"
					return
				}
			}
		}()

		return completed
	}
	cancel := make(chan bool)
	strings := make(chan string)
	terminated := doWork(cancel, strings) // block for ever (receiving value from nil channel)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(cancel)
	}()
	go func() {
		// small leak
		for i := 0; i < 50; i++ {
			time.Sleep(time.Millisecond * 100)
			strings <- fmt.Sprintf("i:%d", i)
		}
	}()
	for v := range terminated {
		fmt.Println(v)
	}

	// Perhaps more work is done here
	fmt.Println("Done.")
	fmt.Println(runtime.NumGoroutine())
}
```

The previous example handles the case for goroutines receiving on a channel nicely, but what if we're dealing with the  reverse situation: a goroutine blocked on attempting to write a value to a channel.

```go
func main() {
	newRandStream := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(ch)
			for {
				ch <- rand.Int()
			}
		}()
		return ch
	}
	ch := newRandStream()
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}
```
You can see from the output that the fmt.Println statement never gets run. After the third iteration of our loop, out goroutine blocks trying to send the next random integer to a channel that is no longer being read from. 

```go
func main() {
	newRandStream := func(cancel chan struct{}) <-chan int {
		ch := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(ch)
			for {
				select {
				case ch <- rand.Int():
				case <-cancel:
					return
				}
			}
		}()
		return ch
	}
	cancel := make(chan struct{})
	ch := newRandStream(cancel)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	cancel <- struct{}{}

	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}
```

**Now that we know how to ensure goroutines don't leak, we can stipulate a convention: If a goroutine is responsible from creating a goroutine, it is also responsible from ensuring it can stop the goroutine**

### The or-channel

The or-channel pattern creates a composite done channel through recursion and goroutines. Let's have a look:

### Error Handling

### Pipelines

A pipeline is just another tool you can use to form an abstraction in your system. In particular, it is very powerful tool to use when your program needs to process stream, or batches of data. A pipeline is nothing more than a series of things that take data in, perform an operation on it, and pass the data back out. We call each of these operations a stage of the pipeline.

By using a pipeline, you separate the concerns of each stage, which provide numerous benefits. You can modify stages independent of one another, you can mix and match how stages are combined independent of modifying the stages, you can process each stage concurrent to upstream or downstream stages.

Here is a function that could be considered a pipeline stage:

```go
multiply := func(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v*multiplier
	}
	return multipliedValues
}
```

```go
add := func(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}
```

Combine them
```go
ints := []int{1,2,3,4}
for _, v := range add(multiply(ints, 2), 1) {
	fmt.Println(v)
} 
```


Without pipeline
```go
ints := []int{1,2,3,4}
for _, v := range ints {
	fmt.Println(2*(v*2+1))
}
```

This looks much simpler, the procedural code doesn't provide the same benefits a pipeline does when dealing with streams of data. Notice how each stage is talking a slice of dat and returning a slice of data? These stages are performing what we call batch processing. This just means that they operate on chunks of data all at once instead of one discrete value at a time. There is another type of pipeline stage that performs stream processing. This means that the stage receives and emits one element at a time.

There are pros and cons to batch processing versus stream processing, which we'll discuss in just a bit. For now, notice that for the original data to remain unaltered, each stage has to make a new slice of equal length to store the results of its calculations. That means that the memory footprint of our program at any one time is double the size of the slice we send into the start of our pipeline.

Let's convert our stages to be stream oriented and see what that looks like:
```go
func main() {
	multiply := func(value int, multiplier int) int {
		return value * multiplier
	}
	add := func(value int, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(add(multiply(v, 2), 1))
	}
}
```
### Best Practices for Constructing Pipelines

Channels are uniquely suited to constructing pipelines in Go because the fulfill all of our basic requirements. They can receive and emit values, they can safety be used concurrently, they can be ranged over, and they are reified by the language. 

```go
func main() {
	generator := func(done <-chan any, integers ...int) <-chan int {
		rs := make(chan int, len(integers))
		go func() {
			defer close(rs)
			for _, v := range integers {
				select {
				case <-done:
					return
				case rs <- v:
				}
			}
		}()
		return rs
	}

	multiply := func(done <-chan any, intStream <-chan int, multiplier int) <-chan int {
		rs := make(chan int, len(intStream))
		go func() {
			defer log.Println("exit multiply stage")
			defer close(rs)
			for {
				select {
				case <-done:
					return
				case v, isOpen := <-intStream:
					if !isOpen {
						return
					}
					rs <- v * multiplier

				}
			}
		}()
		return rs
	}

	add := func(done <-chan any, intStream <-chan int, additive int) <-chan int {
		rs := make(chan int, len(intStream))
		go func() {
			defer log.Println("exit add stage")
			defer close(rs)
			for {
				select {
				case <-done:
					return
				case v, isOpen := <-intStream:
					if !isOpen {
						return
					}
					rs <- v + additive
				}
			}
		}()
		return rs
	}

	done := make(chan any)

	_ = generator
	close(done)
	// intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, nil, 2), 1), 2k)
	for v := range pipeline {
		log.Println(v)
	}
}
```	

```go
func main() {
	generator := func(done <-chan any, integers ...int) <-chan int {
		rs := make(chan int, len(integers))
		go func() {
			defer close(rs)
			for _, v := range integers {
				select {
				// prevent preemptable by closing the done channel
				case <-done:
					return
				case rs <- v:
				}
			}
		}()
		return rs
	}

	multiply := func(done <-chan any, intStream <-chan int, multiplier int) <-chan int {
		rs := make(chan int, len(intStream))
		go func() {
			defer log.Println("exit multiply stage")
			defer close(rs)
			for {
				select {
				// prevent preemptable by closing the done channel
				case <-done:
					return
				case v, isOpen := <-intStream:
					if !isOpen {
						return
					}
					rs <- v * multiplier

				}
			}
		}()
		return rs
	}

	add := func(done <-chan any, intStream <-chan int, additive int) <-chan int {
		rs := make(chan int, len(intStream))
		go func() {
			defer log.Println("exit add stage")
			defer close(rs)
			for {
				select {
				// prevent preemptable by closing the done channel
				case <-done:
					return
				case v, isOpen := <-intStream:
					if !isOpen {
						return
					}
					rs <- v + additive
				}
			}
		}()
		return rs
	}

	done := make(chan any)
	close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		log.Println(v)
	}
}
```

### Some Handy Generators

```go
func main() {
	repeat := func(done <-chan any, value ...any) <-chan any {
		valueStream := make(chan any)
		go func() {
			defer close(valueStream)
			for {
				for _, v := range value {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	take := func(done <-chan any, valueStream <-chan any, num int) <-chan any {
		takeStream := make(chan any)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan any)
	for v := range take(done, repeat(done, 1), 20) {
		fmt.Println(v)
	}
}
```

### Fan-Out, Fan-in

Sometimes, stages in your pipeline can be particularly computationally expensive. When this happens, upstream stages in your pipeline can be blocked while waiting for your expensive stages to complete. Not only that, but the pipeline itself can take a long time to execute as a whole.

Fan-out is a term to describe the process of starting multiple goroutines to handle input from the pipeline, and fan-in is a term to describe the process of combining multiple results into one channel.

So what makes a stage of a pipeline suited for utilizing this pattern? You might consider fanning out one of your stages if both of the following apply:

- It doesn't rely on values that the stage had calculated before.
- It takes a long time to run.

The property of order-independence is important because you have no guarantee in what order concurrent copies of your stage will run, nor in what order they will return.

```go
// fan-out example
```

Fanning in means multiplexing or joining together multiple streams of data into a single stream.
```go
fanin := func(done <-chan any, channels ...<-chan any) <-chan any {
	var wg sync.waitgroup
	multiplexedstream := make(chan any)
	multiplex := func(c <-chan any) {
		defer wg.done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedstream <- i:
			}
		}
	}

	// select from all the channels
	wg.add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.wait()
		close(multiplexedstream)
	}()
}
```

### The or-done-channel

### The tee-channel

Sometimes you may want to split values coming in from a channel so that you can send them off into two separate areas of your codebase. Imagine a channel of user commands: you might want to take in a stream of user commands on a channel, send them to something that executes them, and also send them to something that logs the commands for later auditing.


```go
func main() {
	tee := func(done <-chan struct{}, in <-chan any) (_, _ <-chan any) {
		out1 := make(chan any)
		out2 := make(chan any)
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range in {
				// We will want to use local versions of out1 and out2, so we shadow these variables
				// Notice that writes to out1 and out2 are tightly coupled. The iteration over in cannot continue until both out1 and out2 
				// have been written to. 
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					// We're going to use one select statement so that writes to out1 and out2 don't block each other. 
					// To ensure both are written to, we'll perform two iterations of the select statement
					select {
					case <-done:
					case out1 <- val:
						// Once we've written to a channel, we set its shadowed copy to nil so that further writes will block and the other channel may continue
						// Send value to nil channel always block
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	generator := func(done <-chan struct{}, vals ...int) <-chan any {
		ch := make(chan any)
		go func() {
			defer close(ch)
			for _, v := range vals {
				ch <- v
			}
		}()

		return ch
	}

	done := make(chan struct{})
	out1, out2 := tee(done, generator(done, 1, 2))

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range out1 {
			fmt.Println(v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range out2 {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
```

### The bridge-channel

In some circumstances, you may find yourself wanting to consume values from a sequence of channels:

```go
func main() {
	bridge := func(done <-chan any, chanStream <-chan <-chan any) <-chan any {
		// This is the channel that will return all values from bridge
		valStream := make(chan any)
		go func() {
			defer close(valStream)
			// This loop is responsible for pulling channels off of chanStream and providing them to a nested loop for use
			for {
				var stream <-chan any
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				// This loop is responsible for reading values off the channel it has been given and repeating those values onto valStream. When the stream we're currently
				// looping over is closed, we break out of
				for val := range stream {
					select {
					case <-done:
					case valStream <- val:
					}
				}
			}
		}()
		return valStream
	}
}
```

A sequence of channels suggests an ordered write, albeit from different sources. One example might be a pipeline stage whose lifetime is intermittent.

### Queuing

Sometimes it's useful to begin accepting work for your pipeline even though the pipeline is not yet ready for more. This process is called queuing.
While introducing queuing into your system is very useful, it's usually one of the last techniques you want to employ when optimizing your program.



## Chapter 5: Concurrency at Scale

### Error Propagation

Error needs to relay a few pieces of critical information:

- What happened.
This is the part of the error that contains information about what happened "disk full", "socket closed" or "credentials expired". This information is likely to be generated implicitly by whatever it was that generated the errors.

- When and where it occurred.
Errors should always contain a complete stack trace starting with how the call was initiated and ending with where the error was instantiated. The error should contain the time on the machine the error was instantiated, in UTC.

- A friendly user-facing message.
The message that gets displayed to the user should be customized to suit your system and its users.

- How the user can get more information.
At some point, someone will likely want to know, in detail, what happened when the error occurred.

## Heartbeats

Heartbeats are a way for concurrent processes to signal life to outside parties. They get their name from human anatomy wherein a heartbeat signifies life to an observer. Heartbeats have been around since before Go, and remain useful within it.

There are two different types of heartbeats we'll discuss in this section:
- Heartbeats that occur on a time interval
- Heartbeats that occur at the beginning of a unit of work

## Chapter 6: Goroutines and the Go Runtime

### Work Stealing
Go will handle multiplexing goroutines onto OS threads for you. The algorithm it uses to do this is now as work stealing strategy.
