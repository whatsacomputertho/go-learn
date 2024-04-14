# Defer, panic, and recover

In this application, we explore some more advanced control flow constructs in go.

- [Defer, panic, and recover](#defer-panic-and-recover)
  - [Defer](#defer)
  - [Panic](#panic)
  - [Recover](#recover)

## Defer

Using `defer`, we can invoke a function but delay its execution until the parent function returns.  This is generally used to group "open" and "close" operations together, even though the "close" operation needs to happen after all of the logic.

Here is an example of a function which uses `defer` to close an HTTP connection after its response body has been read into memory.  Using `defer`, we can declare the close statement together with the open statement, but it won't fire until the function exits.
```go
// Practical example of defer in go
func practicalDefer() {
	// GET placeholder JSON from jsonplaceholder.typicode.com
	// Once finished reading the response body, close it
	// This is the most common use case of defer in go
	// It lets us write "open" and "close" logic right
	// next to one another, while still closing at the end
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Won't close until the function returns

	// Read the response body from our GET request
	jsonBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	fmt.Printf("%s\n", jsonBody)
}
```

This should be applied carefully in the case of loops.  Note that the `defer` keyword will not delay execution until the end of the loop's iteration, but rather the end of the function.  So if we were to wrap the above in a loop, we would have a bunch of open HTTP connections until the loop completes and the function exits.

Defer statements run in LIFO (last-in, first-out) order.
```go
/*
Demonstrates the last-in first-out nature of defer in go

Result:
I should print third
I should print second
I should print first
*/
func lifoDefer() {
	defer fmt.Println("I should print first")  // Will print third
	defer fmt.Println("I should print second") // Will print second
	fmt.Println("I should print third")        // Will print first
}
```

Arguments passed into deferred function calls are eagerly evaluated at defer time, not at execution time.
```go
// Example of eager argument resolution using defer
func eagerArgResDefer() {
	a := "start"
	defer fmt.Println(a) // Prints "start", not "end"
	a = "end"
}
```

## Panic

Using `panic`, we can signify that our go program can no longer continue to function and must exit.  The go runtime is capable of `panic`-ing under certain conditions.

In general, the idiomatic way to handle exceptional events (that is, something that would raise an Exception in some other language), is to return an error value along with the function's return value.  Importantly, this means the program will NOT exit, it will continue to execute and it is up to the programmer to decide (sparingly) which events are worthy of a hard exit.

Here is an example of this design pattern in go.  We see the `net/http` module returns two values, a response and an error value.  If the HTTP request times out for example, this will not yield an exit condition.  Instead, it will just populate the error value, and it is up to us to ensure that we always check that value and handle it appropriately.
```go
func errReturn() {
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        log.Fatal(err)
    }
    // ...(function continues below)
}
```

Given the above, the `panic` function can be used to signal to the go runtime that our program has reached an unrecoverable state.  In the case of network programming, we might use `panic` if we can't obtain a TCP port for our web server.

When we do use `panic`, the function in which the `panic` occurred will stop executing.  If we do not handle the `panic`, it will propagate up the call stack and cause the program to exit.  Any `defer`-ed functions will execute before the program exits.
```go
/*
Example of defer and panic used together

Result:
Start
This was deferred
panic: Goodbye, cruel world!

goroutine 1 [running]:
main.deferAndPanic()
        /path/to/my/go/executable/defer-panic-recover/defer-panic-recover.go:139 +0x12
main.main()
        /path/to/my/go/executable/defer-panic-recover/defer-panic-recover.go:75 +0xabc
exit status 2
*/
func deferAndPanic() {
	fmt.Println("Start")
	defer fmt.Println("This was deferred") // Will run
	panic("Goodbye, cruel world!")         // Will run
	fmt.Println("End")                     // Will not run
}
```

## Recover

Using `recover`, we can specify how our program might continue to function even in the instance of a `panic`.  This is particularly useful if some library implements `panic` non-sparingly.  We know we have this tool under our belt to recover from a recoverable panic.
```go
/*
Example of panic and recover used together

Result:
About to panic
2024/04/14 09:15:38 Goodbye, cruel world!
...(Logs continue - program does not exit)
*/
func panicAndRecover() {
	fmt.Println("About to panic")

    // Defer an anonymous function call which recovers from the panic
	defer func() {
        // This recover call will cause the panic to stop propagating
		if err := recover(); err != nil {
			log.Println(err) // Will run
		}
	}()

    // Panic - this will fire the above function before exiting
	panic("Goodbye, cruel world!") // Will run

	fmt.Println("Done panicking")  // Will not run
}
```

A `recover` function call is only useful within a deferred function call.  This is because deferred function calls will run just before its immediate parent function call returns.  In the case of a panic, this is just before the panic starts to propagate up the call stack.

It is also worth noting that `recover` will not cause the immediate parent function to continue to execute.  It will, however, allow function calls higher on the call stack to continue their execution.