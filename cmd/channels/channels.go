package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// Log entry struct used for select & signal-only example
type logEntry struct {
	time     time.Time
	severity string
	message  string
}

const logInfo = "INFO"

// Long-lived global log & done channels
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Signal-only channel

/*
Select & signal-only channels

This function is executed via a goroutine below which runs
throughout the entirety of the program.  We use a select
statement to check for whether we are done logging, and close
the log channel if so.
*/
func logger() {
loggerloop:
	for {
		// This is a "blocking select statement"
		//
		// The code won't execute until a message comes into one
		// of the channels.  If we wanted it to be a "non-blocking
		// select statement", then we would add a default case,
		// which would execute on any iteration in which a message
		// is not received in either of the channels.
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break loggerloop
		}
	}
}

func main() {
	/*
		Channel basics

		Here we begin with a basic example of communication between
		two goroutines.

		A common use case of goroutines and channels together is
		the case in which data is generated and processed in an
		asynchronous manner.  Maybe the data takes a long time to
		generate, but is processed quickly, or vice versa.
	*/
	fmt.Println("#### Channel basics ####")

	// Example 1 - Basic communication between goroutines using channels
	ch := make(chan int)
	wg.Add(2)
	go func() {
		// Extract an int from the channel and print it
		i := <-ch
		fmt.Println(i) // It should print 42
		wg.Done()
	}()
	go func() {
		// Input an int into the channel
		i := 42
		ch <- i // Copies i
		i = 27  // Mutation does not affect channel communication
		wg.Done()
	}()
	wg.Wait()

	// Example 2 - More basic communication between goroutines
	// This time we loop over the send & receive
	for j := 0; j < 3; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 33
			wg.Done()
		}()
	}
	wg.Wait()

	// Example 3 - Deadlock condition
	// Since we are not using a buffered channel, we cannot add
	// another value into it before the current channel value is
	// read.
	//wg.Add(1)
	//go func() {
	//	i := <-ch
	//	fmt.Println(i)
	//	wg.Done()
	//}()
	//for j := 0; j < 3; j++ { // Will fail on third iteration
	//	wg.Add(1)
	//	go func() {
	//		ch <- 42
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	fmt.Println("")

	/*
		Restricting data flow

		When using channels in the way that we have used them so
		far, they can be used for bidirectional read-write.  This
		is sometimes desirable, but not always.  More often, you
		will want to dedicate goroutines as readers and writers.

		The way in which we define read-only and send-only channels
		in go looks like polymorphism.  We take a bidirectional
		channel and have it behave as though it were read-only or
		write-only.  However, this is behavior specific to channels
		that the go runtime supports.  In effect, the channel is
		being cast to a single-directional channel by the go runtime.
	*/
	fmt.Println("#### Restricting data flow ####")

	// Example 1 - Using a channel for bidirectional read-write
	wg.Add(2)
	go func() {
		i := <-ch      // Getting value from other goroutine
		fmt.Println(i) // Should print 42
		ch <- 27       // Sending value to other goroutine
		wg.Done()
	}()
	go func() {
		ch <- 42          // Sending value to other goroutine
		fmt.Println(<-ch) // Should print 27
		wg.Done()
	}()
	wg.Wait()

	// Example 2 - Using read-only and write-only channels
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		//ch <- 27 // Error: Cannot send to receive-only channel
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		//fmt.Println(<-ch) // Error: Cannot receive from send-only channel
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println("")

	/*
		Buffered channels

		We can define channels such that they allocate an
		internal buffer.  This is used generally for cases in
		which the sender and receiver operate on different
		frequencies.
	*/
	fmt.Println("#### Buffered channels ####")

	// Example 1 - Buffered channels
	bufCh := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch      // One receiver
		fmt.Println(i) // Still works, no error
		wg.Done()
	}(bufCh)
	go func(ch chan<- int) {
		ch <- 42 // Two senders
		ch <- 27 // This value is lost though
		wg.Done()
	}(bufCh)
	wg.Wait()
	fmt.Println("")

	/*
		For loops with channels

		How do we deal with senders which send more values than
		we have receivers?  We can use a for-range loop over the
		channel.

		We first attempt to do so naively by looping over the
		receiver channel, but we reach a deadlock condition.  The
		channel can store any number of values, so we continue
		monitoring indefinitely even after all messages are sent.
		This leads to a deadlock condition.

		The solution for this is to close the channel in the sender
		goroutine after sending the final message.  This signals to
		the receiver that all messages have been sent.

		We also note that we may close channels, channels cannot be
		reopened, and sending values into a closed channel leads to
		a runtime panic.

		We finally note that the for-range loop is syntactic sugar
		for what can be achieved explicitly via a generic for loop.
		We show that the act of reading from a channel returns an
		ok boolean signifying whether the channel is closed. We
		remark that this is useful for instances in which we need
		to process channel data outside a loop.
	*/
	fmt.Println("#### For loops with channels ####")

	// Example 1 - For range loop over channel, deadlock condition
	//loopCh1 := make(chan int, 50)
	//wg.Add(2)
	//go func(ch <-chan int) {
	//	for i := range ch { // However here we reach a deadlock condition
	//		fmt.Println(i) // Continue monitoring for messages when no more
	//	}
	//	wg.Done()
	//}(loopCh1)
	//go func(ch chan<- int) {
	//	ch <- 42 // Two senders
	//	ch <- 27 // This value is not lost this time
	//	wg.Done()
	//}(loopCh1)
	//wg.Wait()

	// Example 2 - For range loop over channel, no deadlock
	loopCh2 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(loopCh2)
	go func(ch chan<- int) {
		ch <- 42  // Two senders
		ch <- 27  // This value is not lost this time
		close(ch) // Closing channel signals to receiver that all
		wg.Done() // messages are sent
	}(loopCh2)
	wg.Wait()

	// Example 3 - For range loop over channel, premature closure
	//loopCh3 := make(chan int, 50)
	//wg.Add(2)
	//go func(ch <-chan int) {
	//	for i := range ch {
	//		fmt.Println(i)
	//	}
	//	wg.Done()
	//}(loopCh3)
	//go func(ch chan<- int) {
	//	ch <- 42
	//	close(ch) // Close channel prematurely
	//	ch <- 27  // Go runtime panics here
	//	wg.Done() // with "send on closed channel"
	//}(loopCh3)
	//wg.Wait()

	// Example 4 - Explicit loop over channel using comma-ok syntax
	loopCh4 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			// Channel read returns ok boolean
			// Break if ok boolean is false - channel was closed
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(loopCh4)
	go func(ch chan<- int) {
		ch <- 42  // Two senders
		ch <- 27  // This value is not lost this time
		close(ch) // Closing channel signals to receiver that all
		wg.Done() // messages are sent
	}(loopCh4)
	wg.Wait()
	fmt.Println("")

	/*
		Select & signal-only channels

		We might encounter a situation where we would like to keep
		a channel open for the entire duration of a program, only
		to close it at the very end of the program's execution.

		Naively one might defer an anonymous function in the main
		function to close the global channel at the end of the main
		function execution, and this is okay in most cases.

		However, we can also use a select statement to deal with
		long-lived channel closure.  We demonstrate a pattern for
		using a select statement together with what is known as a
		signal-only channel to ensure safe closure of a long-lived
		channel.

		A signal-only channel is a channel which accepts an empty
		struct.  This requires no memory allocation, and only acts
		as a flag for whether a message was sent or not.
	*/
	fmt.Println("#### Select & signal-only channels ####")
	go logger()        // Start our logger goroutine
	logCh <- logEntry{ // Send a log message
		time.Now(),
		logInfo,
		"Starting application",
	}
	logCh <- logEntry{ // Send another log message
		time.Now(),
		logInfo,
		"Finishing application",
	}
	time.Sleep(100 * time.Millisecond)
	// Signal completion via the signal-only channel
	// This will close the logger goroutine
	doneCh <- struct{}{}
}
