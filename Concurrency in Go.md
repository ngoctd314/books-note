# Concurrency in Go

## CHAPTER 1. An Introduction to Concurrency

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