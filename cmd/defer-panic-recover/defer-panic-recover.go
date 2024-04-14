package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	/*
		Defer

		The defer keyword allows us to execute a function
		call just before its parent function returns.  Defer
		functions are "LIFO", so the last deferred function
		call is the first function call made before the function
		returns.

		Defer is commonly used when "open" and "close" function
		calls are required to read some resource, like an API
		call or reading from a file.

		Arguments passed into deferred function calls are
		eagerly evaluated, meaning their value(s) at defer time
		are what are ultimately passed into the function call
		at return time.
	*/
	fmt.Println("#### Defer ####")
	basicDefer()
	practicalDefer()
	eagerArgResDefer()
	fmt.Println("")

	/*
		Panic

		In go, errors are handled differently than other
		languages.  It is idiomatic in go to return an error
		return value along with the return value of a function
		to signify if an error occurred.  Then it is up to the
		programmer to decide what to do with that.

		As implied by the above, there are no exceptions in go.
		In their place, we have what is called panic.  This is
		a way to signal that the go program has reached a point
		where it cannot continue.

		It is commonplace to write explicit panicking logic since
		rarely in go modules do we ever panic.  Instead we return
		an error, and allow the programmer to panic if they think
		this is necessary.

		It is worth noting that panics happen after deferred
		statements.  This is important as if we hit a panic, we
		do not need to worry about open resources remaining open
		following the panic.  If we properly defer their close
		statement, they will still close.
	*/
	fmt.Println("#### Panic ####")

	// This will cause the go runtime to panic - dividing by
	// zero is syntactically correct but it cannot be evaluated
	fmt.Println("Dividing by zero causes the go runtime to panic")
	//a, b := 1, 0
	//ans := a / b
	//fmt.Println(ans)

	// Example of custom panicking behavior
	fmt.Println("We can explicitly panic in go")
	//panic("Goodbye, cruel world!")

	// Example of using defer and panic together
	fmt.Println("Deferred statements are executed before panicking")
	//deferAndPanic()
	fmt.Println("")

	/*
		Recover

		We can explicitly recover from a panic using the recover
		function.  Here we demonstrate an example of recovering
		from a panic.

		We see that recovering from a panic will still result in
		the function call where the panic occurred exiting early.
		However, it will not propagate further up the call stack,
		so execution will continue at that higher level.
	*/
	fmt.Println("#### Recover ####")

	// Example of panicking and recovering from it
	fmt.Println("Start") // Will run
	panicAndRecover()    // Will panic and recover from it
	fmt.Println("End")   // Will run
}

// Basic example of defer in go
func basicDefer() {
	defer fmt.Println("I should print first")  // Will print third
	defer fmt.Println("I should print second") // Will print second
	fmt.Println("I should print third")        // Will print first
}

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

// Example of eager argument resolution using defer
func eagerArgResDefer() {
	a := "start"
	defer fmt.Println(a) // Prints "start", not "end"
	a = "end"
}

// Example of defer and panic used together
func deferAndPanic() {
	fmt.Println("Start")
	defer fmt.Println("This was deferred") // Will run
	panic("Goodbye, cruel world!")         // Will run
	fmt.Println("End")                     // Will not run
}

// This funcion panics, but defers an anonymous function
// call which recovers from the defer
func panicAndRecover() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err) // Will run
		}
	}()
	panic("Goodbye, cruel world!") // Will run
	fmt.Println("Done panicking")  // Will not run
}
