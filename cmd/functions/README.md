# Functions

Here we cover the usage of functions in go.

- [Functions](#functions)
  - [Basic syntax](#basic-syntax)
  - [Parameters](#parameters)
  - [Return values](#return-values)
  - [Anonymous functions \& functions as types](#anonymous-functions--functions-as-types)
  - [Methods](#methods)

## Basic syntax

Functions are identified by the `func` keyword.

The function name follows, and the casing of the function name drives visibility in the same way as is done for variables.

The function name is followed immediately by a pair of parenthesis.  In this case, the parenthesis are empty, but in the future we will explore cases in which parameters are defined within the parenthesis.

Then, the function body is defined within the curly braces following the function name and parenthesis. The go compiler enforces that the opening curly brace is included on the same line as the function signature.

Here is an example of a basic function in go.
```go
func helloWorld() {
	fmt.Println("Hello, world!")
}
```

## Parameters

Parameters are values which can be passed into the scope of the function, and which can influence the manner in which the function behaves.

In the function signature's parenthesis, parameters are defined as a comma-separated list of parameter names and types within the.
```go
func sayHello(name string) {
    fmt.Println("Hello", name)
}
```

Parameters of the same type can be listed as a comma separated list followed by their type name at the end of the list.
```go
func sayGreeting(greeting, name string) {
    fmt.Println(greeting, name)
}
```

Functions cannot mutate ordinary parameters beyond their own isolated inner scope.
```go
func changeName(name string) {
    name = "New name" // Has no effect outside the function
}
```

However, when parameters are passed in as a pointer, then functions can mutate their values globally.
```go
func changeName(name *string) {
    *name = "New name" // Updates the parameter value globally
}
```

We can accept what are called variadic parameters in our function signature in the following way.  These parameters allow for any number of values of that same type to be passed in.  It must be the last parameter in the function signature, and it is stored as a slice within the function's scope.
```go
func sumManyAndPrint(values ...int) {
    // Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

    // Print the result
    fmt.Println(result)
}
```

## Return values

Return values are values which can be passed back into the parent scope from inside the function.  In order to specify a return value, we must include its type in the function signature just following the parenthesis.  Then we need to explicitly return it when we're ready to do so using the `return` keyword.  Once we reach a `return`, the function stops executing at that point (except for any deferred functions).
```go
func sumManyAndReturn(values ...int) int {
	// Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

	// Return the result
	return result
}
```

We can produce multiple return values from a function.  To do so, the function signature must list the return types in parenthesis.  To exemplify this, we use one of the most common idioms in go, which involves returning a return value together with an error value to indicate whether an error occurred during the execution of the function.
```go
func divideFloats(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
```

We can also use named return values.  In this case, we define a name for the return type in parenthesis in the function signature.  That type is instantiated into a variable of that name within the function's scope implicitly.  It can be mutated within the function scope, and then returned implicitly by just writing `return` with no explicit following value.
```go
func sumManyAndReturnImplicitly(values ...int) (result int) {
	// Loop through the values and sum into result
	for _, v := range values {
		result += v
	}
	return // Returns result variable
}
```

Lastly, we can return pointers to values defined within the function.  Under the hood, the go runtime promotes the value to a common location on the heap, and out of the local function stack.
```go
func sumManyAndReturnRef(values ...int) *int {
	// Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

	// Return a pointer to the result
	return &result
}
```

## Anonymous functions & functions as types

Anonymous functions are functions which are not named.  They can be immediately invoked like so, which is not commonly used but may be useful for creating an isolated scope within a function.
```go
func() {
    fmt.Println("Hello, world")
}() // Hello, world
```

They can also be assigned as a variable, or passed as an argument into another function call.
```go
f := func() {
    fmt.Println("Hello, world")
}
f() // Hello, world
```

We therefore start to see that functions themselves are treated as types in go.  They can be assigned to variables, passed into other functions, and returned by functions.  The function type signature is similar to the regular function signature explored above, but with no parameter names given.
```go
var f func(string, string, int) (int, error)
```

## Methods

Methods are special functions which execute in the context of a type.  Methods are commonly associated with structs, but can be associated with any type.  We can associate a method with a type in the following way, including a named type between the `func` keyword and the function name.
```go
// Type definition for greeter
type greeter struct {
	greeting string
	name     string
}

// Greet method attached to greeter type
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
```

To call this function, we would first instantiate a `greeter` type, and then call it using the dot operator.
```go
// Initialize a greeter struct instance
myGreeter := greeter{
    name:     "Dave",
    greeting: "Hello",
}

// Call the greet function
myGreeter.greet()
```

However, since we defined the receiver as `(g greeter)` in the method signature we note that this will copy the `myGreeter` instance each method call.  This is called a **value receiver**.

To avoid this, we can alternatively use pointers to only pass a reference to the underlying type instance into the function.  This is appropriately called a **pointer receiver**.
```go
func (g *greeter) greetRef() {
	fmt.Println(g.greeting, g.name)
}
```