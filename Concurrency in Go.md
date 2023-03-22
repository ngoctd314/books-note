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
