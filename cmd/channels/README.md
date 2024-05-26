# Channels

In our last module we learned that Go abstracts the concept of a thread into a higher-level structure called a goroutine.  We ended on a somewhat open-ended note as we could not identify a perfectly consistent way to handle communication between threads in a concurrent application.

Here, we explore `channels`, which are another higher-level construct native to go which enable safe communication across goroutines.

- [Channels](#channels)
  - [Channel basics](#channel-basics)
  - [Restricting data flow](#restricting-data-flow)
  - [Buffered channels](#buffered-channels)
  - [For loops with channels](#for-loops-with-channels)
  - [Select statements](#select-statements)

## Channel basics

We can create a channel using the built-in `make` function, and this is the only way to create a channel in go.  These channels are strongly typed, thus only a certain type may be sent over a channel.
```go
ch := make(chan int)
```

We can send a message into a channel using arrow syntax, and we can read a value from a channel using similar arrow syntax.
```go
val := 27
ch <- val            // Send val into channel
readVal := <-ch      // Read val from channel
fmt.Println(readVal) // 27
```

One can have multiple senders and receivers of a channel, however one must be careful of some of the error conditions around imbalanced channel writes and reads.

## Restricting data flow

By default, channels are bidirectional, thus values can be sent into and read from a channel.  However, we can cast a channel into a send-only or receive-only channel.  The go runtime then enforces the direction of the channel.
```go
sendOnlyCh := make(chan<- int)
readOnlyCh := make(<-chan int)
```

## Buffered channels

By default, channels block the sender side until a receiver is available.  Similarly, channels block the receiver side until a message is available.  Thus, the number of receivers must equal the number of senders.  If not, the result is a deadlock condition.

Buffered channels are one tool which allow us to deal with senders and receivers which operate on different frequencies.  Suppose the sender sends its values in bursts; then a buffered channel will allow us to accept a predetermined number of these values without them being immediately read.
```go
bufCh := make(chan int, 50) // Internal buffer of 50 int values
```

## For loops with channels

For-range loops over channels allow us to receive values from a channel in a loop construct.  This is another effective tool for dealing with senders and receivers which operate on different or varying frequencies.
```go
// Here, ch is a channel
for i := range ch {
  fmt.Println(i)
}
```

Note that the for-range loop will continuously monitor the channel for messages until it closes.  Also note that here, we do not accept an index as a return value from the `range` keyword, but instead we receive the channel value only.

Closing a channel is achieved by using the built-in `close` function.  This is a one-time action which prohibits any further messages from being sent into the channel.  Channels may not be reopened after closing.
```go
close(ch) // Here, ch is a channel
```

## Select statements

Select statements allow a goroutine to monitor multiple channels at once.  It blocks if all channels are blocked, and if multiple channels receive values at once, then the behavior is undefined.  There is no guarantee which value will be processed in that case.
```go
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Signal-only channel

// Logger example using a log channel and a "done" signal-only channel
// Breaks if "done" signal-only channel is signaled
func logger() {
loggerloop:
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break loggerloop
		}
	}
}
```