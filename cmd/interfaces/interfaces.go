package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
Basics of interfaces

Interfaces are types in go, and to define our own interfaces
we can simply define a custom type assigned to an interface
definition.  Interfaces simply store method signatures.  If
some other type implements that interface, then we know that
we can freely call the functionality of the interface using
that type.

Interfaces follow the same naming convention as variables and
functions in go, with the added constraint that the interface
name should be the name of one of its methods, plus "er".  In
this case, we're creating an interface called "Writer", and
its main function is appropriately called "Write".
*/
type Writer interface {
	Write([]byte) (int, error)
}

/*
Basics of interfaces

Here we define a ConsoleWriter struct, which is an empty struct
that implements the Writer interface.  Notice that we implement
the interface implicitly by simply defining methods on the
ConsoleWriter struct which share the same signature as the
Writer interface.
*/
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
Basics of interfaces

Interfaces can be defined on any type, not just structs.  Here
we show an example of an interface which we call "Incrementer"
defined on an int type alias which we call "IntCounter".
*/
type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/*
Interface composition

Here we define a new interface, "Closer", and we use it
together with the existing "Writer" interface in a composition
of multiple interfaces.  The main thing we want to show here
is that, in order to implement a composed interface like the
WriterCloser interface defined below, we simply need to
implicitly implement all of its methods.
*/
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	// Write the bytes to the internal buffer
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	// If the buffer has at least 8 characters, then it will
	// write to the console
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	// Flush the buffer
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	/*
		Basics of interfaces

		Here we initialize an instance of the ConsoleWriter
		struct from above, and we write to the console using
		its Write method from the Writer interface.

		We also do so with the Incrementer interface to show
		that interfaces can be leveraged against any type, not
		just structs.
	*/
	fmt.Println("#### Basics of interfaces ####")

	// Example of defining a struct as an interface instance
	var w Writer = ConsoleWriter{}

	// Example of calling an interface method on that struct
	w.Write([]byte("Hello, interfaces!"))

	// Example of defining a primitive as an interface instance
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 3; i++ {
		// Example of calling an interface method on a primitive
		fmt.Println(inc.Increment())
	}
	fmt.Println("")

	/*
		Interface Composition

		Here we initialize an instance of the BufferedWriterCloser
		interface to exemplify the usage of composed interfaces.
	*/
	fmt.Println("#### Interface composition ####")

	// Initialize a WriterCloser
	var wc WriterCloser = NewBufferedWriterCloser()

	// Call its Writer & Closer interface methods
	// These are composed in the WriterCloser interface
	wc.Write([]byte("Hello, interface composition!"))
	wc.Close()
	fmt.Println("")

	/*
		Type conversion

		Here we show that we can convert an interface instance to
		its original type so long as they match.  We do so by
		converting our WriterCloser instance from above to a
		BufferedWriterCloser instance.

		We also show that the go runtime will panic when we try to
		convert an interface instance to a type that does not
		implement that interface.  However, we show that we can use
		comma-ok syntax to check if our conversion was successful
		or failed.

		We then show usage of the empty interface in go.
	*/
	fmt.Println("#### Type conversion ####")

	// Converting our above WriterCloser to a BufferedWriterCloser
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// Attempting to convert our WriterCloser to an io.Reader
	//bwc = wc.(io.Reader) // This will lead to panic
	//fmt.Println(bwc)

	// Safer attempt to convert WriterCloser to an io.Reader
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// Using the empty interface as a middle-man in a safe
	// type conversion
	var myObj interface{} = NewBufferedWriterCloser()
	if newWc, ok := myObj.(WriterCloser); ok {
		newWc.Write([]byte("Hello, interface type conversion!"))
		newWc.Close()
	}
	fmt.Println("")

	/*
		Type switching and interfaces

		Here we revisit the concept of a type switch in go,
		applying our understanding of interfaces this time.

		We also revisit the notion of method sets, and learn
		that for values, only the value-receiver functions in
		an interface apply.  Meanwhile for pointers to values,
		the pointer-receiver AND value-receiver functions
		apply.
	*/
	fmt.Println("#### Type switching and interfaces ####")

	// Example of type switch to check type of interface instance
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}

	// Attempting to initialize WriterCloser as value
	//var myWc WriterCloser = BufferedWriterCloser{
	//	buffer: bytes.NewBuffer([]byte{}),
	//} // Will fail due to pointer receiver implementation
}
