# Interfaces

Here we explore the usage of interfaces in go.  Interfaces are one of the reasons that Go is as scalable and maintainable as it is, establishing it as a top language today.

- [Interfaces](#interfaces)
  - [Basics of interfaces](#basics-of-interfaces)
  - [Composing interfaces](#composing-interfaces)
  - [Type conversion](#type-conversion)
    - [The empty interface](#the-empty-interface)
    - [Type switches](#type-switches)
  - [Implementing interfaces](#implementing-interfaces)
    - [With values](#with-values)
    - [With pointers](#with-pointers)
  - [Best practices](#best-practices)

## Basics of interfaces

We define and implement interfaces in go like what is shown in the following example.  Here we define the `Writer` interface, and implement the `Writer` interface for the `ConsoleWriter` struct.  We do not need to explicitly implement the interface, the go runtime handles this in the background, enabling implicit interface implementation.
```go
// Writer interface definition
type Writer interface {
  Write([]byte) (int, error)
}

// ConsoleWriter struct will implement Writer interface
type ConsoleWriter struct {}

// Write method definition - this means the ConsoleWriter
// struct implicitly implements the Writer interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
  n, err := fmt.Println(string(data))
  return n, err
}
```

## Composing interfaces

Similar to struct embedding, we can also embed, or **compose** interfaces together.  This is another reason to suggest that many, small interfaces is the ideal approach to interface implementation.  We can define small interfaces as building blocks for pluggable interface compositions.
```go
// Writer interface definition
type Writer interface {
  Write([]byte) (int, error)
}

// Closer interface definition
type Closer interface {
  Close() error
}

// WriterCloser interface composition
type WriterCloser interface {
  Writer
  Closer
}
```

## Type conversion

If we are positive of the underlying type of an interface instance, then we can drill into the interface, extract the type instance, and work with its data directly.
```go
// Initialize a WriterCloser instance, which under the hood is a
// pointer to the BufferedWriterCloser struct, which implements the
// WriterCloser interface.
var wc WriterCloser = NewBufferedWriterCloser()

// Here we convert our WriterCloser instance, which we know to be
// a BufferedWriterCloser pointer, back into a BufferedWriterCloser.
// We do so by using the dot operator followed by the underlying
// type, and we also dereference it in the process since it is a
// pointer to an instance of that type.
bwc := wc.(*BufferedWriterCloser) // Note: panics on fail, use comma-ok here
```

### The empty interface

The empty interface is exactly what it sounds like, it's the only interface which defines no methods.
```go
var i interface{} = 0
```

**Every type in go implements the empty interface**, and thus any type can be initialized as an instance of the empty interface.

### Type switches

[The empty interface](#the-empty-interface) explored above is commonly paired with what is known as a **type switch**.  It is a special form of a `switch` statement whose cases are types, and thus conditional logic on the basis of type can be performed.
```go
var i interface{} = 0
switch i.(type) {
  case int:
    fmt.Println("i is an integer")
  case string:
    fmt.Println("i is a string")
  default:
    fmt.Println("I don't know what i is")
}
```

## Implementing interfaces

There are differences in how interfaces work when considering value types versus pointers.  In particular, the differences lie in how the **method set** of value types are calculated, versus the method set of pointers to values.

### With values

The method set for value instances are **all methods with a value receiver**.

Thus, if we have an interface implementation with a mix of value receivers and pointer receivers, then a value type instance WILL NOT implement that interface.  If we attempt to treat it as such, the go runtime will panic.
```go
// BufferedWriterCloser struct definition
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// Write method implementation from Writer interface
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	// ... (Implementation here)
	return 0, nil
}

// Close method implementation from Closer interface
func (bwc *BufferedWriterCloser) Close() error {
	// ... (Implementation here)
	return nil
}

// Attempt to initialize BufferedWriterCloser value as WriterCloser
// Will fail due to pointer receiver implementation
var myWc WriterCloser = BufferedWriterCloser{
	buffer: bytes.NewBuffer([]byte{}),
}
```

### With pointers

The method set for pointer instances is **the union of all value and pointer receiver methods**.

Thus, if we have an interface implementation with a mix of value receivers and pointer receivers, then a pointer type instance WILL implement that interface.
```go
// Simply use the address-of operator to resolve the above
var myWc WriterCloser = &BufferedWriterCloser{
	buffer: bytes.NewBuffer([]byte{}),
}
```

## Best practices

The go community has developed the following best practices for interface implementation.

**Use many, small interfaces**: Single method interfaces are some of the most powerful & flexible interfaces in go.  These include, for example
- `io.Writer`
- `io.Reader`
- `interface{}` (The empty interface - zero methods)

**Don't export interfaces for types that will be consumed**, but **do export interfaces for types that you are using**:  In other words, interfaces are ideally left to the consumer to define and implement.  This is because go supports implicit interface implementation, as long as a struct implements the same method signature(s) as defined by any interface, then it implements that interface implicitly.

**Design functions and methods to receive interfaces whenever possible**: If you need to access the underlying data for your input type, then it's okay to receive the types themselves.  But in many cases we expect to receive "behavior providers" as input types, in which case interfaces should be received.