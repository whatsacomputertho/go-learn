package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	/*
		GoRoutines

		The construction of OS threads is generally expensive.
		OS threads have their own dedicated function call stacks,
		and are allocated ~1MB of RAM.

		Go exposes an abstraction atop OS threads which it calls
		GoRoutines.  The go runtime has a scheduler which maps Go-
		Routines onto OS threads for a set amount of time.  Go-
		Routines are very inexpensive as a result, it is not un-
		common to see 10,000 to 100,000 GoRoutines running at a time
		in Go applications.
	*/
	fmt.Println("#### GoRoutines ####")

	// Example 1 - Call sayHello in a GoRoutine and wait
	// Spawn a new GoRoutine, then finish
	// Bad practice, but wait a bit for the GoRoutine to finish
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// Example 2 - Anonymous function which uses parent scope vars
	// Spawn a new GoRoutine which executes an anonymous funciton
	// The anonymous function uses a varaible from the main func
	// The anonymous funciton is able to access the variable.
	// This is thanks to closure in the go runtime.
	// Bad practice, dependency between thread & parent scope var
	// Go runtime will continue main thread while goroutine runs
	var msg = "Hello"
	go func() {
		fmt.Println(msg) // Will likely print Goodbye, race condition
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)

	// Example 3 - Resolving the above by passing as parameter
	// This invokes a copy, so the msg param will not mutate when
	// the parent scope mutates it.
	var msg2 = "Hello"
	go func(msg string) {
		fmt.Println(msg) // Will print Hello due to copy
	}(msg2)
	msg2 = "Goodbye"
	time.Sleep(100 * time.Millisecond)
	fmt.Println("")

	/*
		WaitGroups

		A WaitGroup in go is used to syncrhonize GoRoutines.
		Above, we use time.sleep quite a bit to wait for our
		GoRoutines to run.  WaitGroups are abstractions which
		allow us to wait for GoRoutines to run, and signal
		completion to the parent.
	*/
	fmt.Println("#### WaitGroups ####")

	// Example 1 - Applying a WaitGroup to the above example
	// No longer need to guess execution time using time.Sleep
	var msg3 = "Hello"
	wg.Add(1)
	go func() {
		fmt.Println(msg3) // Will print Hello due to Wait & Done
		wg.Done()
	}()
	wg.Wait()
	msg2 = "Goodbye"

	// Example 2 - Displaying inconsistent behavior of concurrency
	// We should expect to see various values printed out, not
	// the standard "(1) Hi, world", "(2) Hi, world" ...
	// Another example of a race condition
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait()
	fmt.Println("")

	/*
		Mutexes

		Mutexes are an abstraction atop global variables used in
		concurrent programming which make it so that parallel
		processes can access data in a consistent manner, one at
		a time.
	*/
	fmt.Println("#### Mutexes ####")

	// Example 1 - Applying a mutex to the above example
	// We should expect to see more consistent behavior since
	// the mutex gates GoRoutines from mutating the global variable.
	// Still a problem here - this is just single-threading with
	// extra steps.  Using the mutex locks globally like so removes
	// any benefit we would have seen from multithreading.
	counter = 0
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go sayHiMutex()
		m.Lock()
		go incrementMutex()
	}
	wg.Wait()
	fmt.Println("")

	/*
		GOMAXPROCS

		Here we briefly cover the GOMAXPROCS runtime API.  It can
		be used to set the max GoRoutines which can be spawned at
		a time.  This is something that should be fine-tuned as
		there are bottlenecks at both ends.  On one hand, running
		with GOMAXPROCS set to 1 will effectively single-thread our
		application.  On the other hand, running with GOMAXPROCS set
		to 100 will lead to an overworked scheduler that outweighs
		any benefit we might have seen from multithreading.

		It is best practice to run your application through a perf
		test suite to find the best value of GOMAXPROCS.
	*/
	fmt.Println("#### GOMAXPROCS ####")

	// Simply read the current max processes
	fmt.Printf("GOMAXPROCS: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Println("Hello, world")
}

/*
WaitGroups

Used in WaitGroups example 2, these functions are written
to be executed as GoRoutines, and they refer to a global
WaitGroup and a global int variable which they both mutate
and read.
*/
func sayHi() {
	fmt.Printf("(%v) Hi, world\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

/*
Mutexes

Used in Mutexes example 1, these functions are written to be
executed as GoRoutines, and they refer to a global WaitGroup
and a global RWMutex which is used to gate access to a global
int variable which they both mutate and read.

Here, the mutex is locked globally, before it is unlocked by the
functions when executed in a GoRoutine.  We note this is bad
practice in the main function.
*/
func sayHiMutex() {
	fmt.Printf("(%v) Hi, mutex\n", counter)
	m.RUnlock()
	wg.Done()
}

func incrementMutex() {
	counter++
	m.Unlock()
	wg.Done()
}
