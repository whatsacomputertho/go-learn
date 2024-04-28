# GoRoutines

Here we cover GoRoutines in go.  GoRoutines are one of the most hyped features in go, as it is a part of what makes go a language which supports concurrency natively, as a first-class citizen.

- [GoRoutines](#goroutines)
  - [Creating GoRoutines](#creating-goroutines)
  - [Synchronization](#synchronization)
    - [WaitGroups](#waitgroups)
    - [Mutexes](#mutexes)
  - [Parallelism](#parallelism)
  - [Best practices](#best-practices)

## Creating GoRoutines

To create a GoRoutine, simply use the `go` keyword in front of a function call.
```go
// Calling a sayHello function as a goroutine
go sayHello()
```

The go runtime implements closure for us so that we can refer to parent scope variables from anonymous functions in goroutines like what is seen below.
```go
// (Race condition) Calling an anonymous function as a goroutine
var msg = "Hello" // Initialize msg variable
go func() {
    fmt.Println(msg) // Will likely print Goodbye, race condition
}()
msg = "Goodbye" // Mutate msg variable
```

However to avoid race conditions in situations like the above, we might pass values as parameters into the funciton call to invoke a copy.
```go
// (No race condition) Calling an anonymous function as a goroutine
var msg = "Hello" // Initialize msg variable
go func(msg string) {
    fmt.Println(msg) // Will print Hello
}(msg) // Pass into anonymous function, invoking a copy
msg = "Goodbye" // Mutate msg variable, does not affect copy
time.Sleep(100 * time.Millisecond)
```

## Synchronization

GoRoutines are very useful for cases in which global control flow and state does not need to be maintained.  We just allow the processes to spawn and run.  But we will very quickly run into issues when we need to store global state, or control the execution of asynchronous processes synchronously from the top-down.

### WaitGroups

The `sync.WaitGroup` abstraction allows us to wait for goroutines until a predetermined number of them complete.  It exposes a few important functions
- `Add(n int)` - Adds n onto the total number of goroutines we need to complete before proceeding
- `Wait()` - Waits for the specified number of goroutines to complete
- `Done()` - Executed from inside a goroutine to signal completion

### Mutexes

The `sync.Mutex` and `sync.RWMutex` abstractions allow us to protect global variables in multithreaded applications, and ensure that they are read and mutated in a consistent manner.

We explored the `sync.RWMutex` in our example, and it exposed the following functions which we used.
- `RLock()` - Locks the underlying data from being read
- `RUnlock()` - Unlocks the underlying data from being read
- `Lock()` - Locks the underlying data from being mutated
- `Unlock()` - Unlocks the underlying data from being mutated

## Parallelism

By default, Go will use CPU threads equivalent to the available number of cores on whatever machine it's running on.  We can fine-tune the number of CPU threads used by our application by setting `runtime.GOMAXPROCS`.  More threads can increase performance, but too many threads can slow it down.
- Day one: Use `GOMAXPROCS > 1`
- Go live: Fine-tune `GOMAXPROCS` according to performance test results

## Best practices

GoRoutines are very powerful, but they do tend to get messy and fall out of hand.  Here are some best practices which we can apply to avoid having our GoRoutines fall out of hand.

**Don't create goroutines in libraries**: Let the consumer control concurrency, do not cause the consumer to be bound by your preferred approach to concurrency.  This advice can be softened a bit if we have function calls which return a `channel` with a result.

**When creating a goroutine, know how it will end**: It's really easy to have a goroutine run indefinitely and fail to be cleaned up.  For example, we might launch a goroutine as a listener process which runs indefinitely.  But this can cause subtle memory leaks which can drain application resources and cause crashes.

**Check for race conditions at compile time**: Add the `-race` flag to `go run`/`go build` to check for race conditions at compile time.  This will cause our compiler to print data race warnings.