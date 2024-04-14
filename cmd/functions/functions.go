package main

import (
	"fmt"
)

/*
Methods

We can define a special type of function called a method.
Methods are associated with type instances.  If I call a
method, I have to have an instance of the type and then
call the function using the dot operator on that type.
*/
type greeter struct {
	greeting string
	name     string
}

/*
Methods

This is the greet method associated with the greeter type.
We specify that this is a method by including in parenthesis
a type and a variable name between the func keyword and the
function name.
*/
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

/*
Methods

Optionally, we can use refer to the type via a pointer reference
so that it is not copied each time we call its associated method.
This is ideal for large types.
*/
func (g *greeter) greetRef() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	/*
		Basic syntax

		Here, we execute a basic hello world function
		which is explored below in an effort to exemplify
		basic function syntax.
	*/
	fmt.Println("#### Basic syntax ####")

	// Example of a basic function call
	helloWorld()
	fmt.Println("")

	/*
		Parameters

		We cannot directly influence the logic performed by a
		function externally.  However, a function may accept
		parameters which it uses internally to influence its
		control flow.

		We show that function parameters are copied when
		provided to a function as arguments.  However, we can
		allow for functions to mutate underlying data if we
		instead provide them with a pointer which references
		data in memory.

		Passing pointers into functions is preferred when
		considering efficiency.  If we intend to pass complex
		data structures into functions, we had better reference
		it via a pointer rather than allowing it to be copied
		just to pass it into the function's scope.

		We can define variatic parameters to pass any number
		of like-typed parameters into a function.
	*/
	fmt.Println("#### Parameters ####")

	// Example of a function which accepts a parameter
	sayHello("Go")

	// Example of a function which accepts multiple parameters
	sayHelloMany("Alice", 3)

	// Example of a function which accepts multiple like-typed params
	sayGreeting("Hello", "Bob")

	// Example of a function which mutates one of its input params
	myGreeting := "Hello"
	myName := "Charlie"
	sayAndMutateGreeting(myGreeting, myName) // Mutates greeting to "Goodbye"
	fmt.Println(myGreeting)                  // Hello

	// Example of a function which mutates one of its input params
	// But this time it uses a pointer to do so
	sayAndMutateGreetingRef(&myGreeting, myName) // Mutates greeting to "Goodbye"
	fmt.Println(myGreeting)                      // Goodbye

	// Example of a function which takes a variatic parameter
	sumMany(1, 2, 3, 4, 5)
	fmt.Println("")

	/*
		Return values

		We can also use functions to construct resultant values
		for us, and return them into the parent scope from which
		we called the function.
	*/
	fmt.Println("#### Return values ####")

	// Example of a function which returns a value
	sum := sumManyAndReturn(1, 2, 3, 4, 5, 6)
	fmt.Println("Returned sum", sum)

	// Example of a function which returns a pointer
	sumRef := sumManyAndReturnRef(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Returned sum", *sumRef)

	// Example of a function which returns a named return value
	sumNamed := sumManyAndReturnImplicitly(1, 2, 3, 4)
	fmt.Println("Returned sum", sumNamed)

	// Example of a function which returns an error along with its return value
	d, err := divideFloats(5.0, 0.0)
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
	fmt.Println("Result", d)
	fmt.Println("")

	/*
		Functions as types

		Here we see that functions are first-class citizens in go,
		and thus they can be treated as types.  It is worth noting
		that obviously functions, when treated as variables, cannot
		be executed before they are defined.
	*/
	fmt.Println("#### Functions as types ####")

	// Here we define and immediately execute an anonymous function
	func() {
		fmt.Println("Hello, world!")
	}()

	// Here we define another one and assign it to a variable
	myFunc := func() {
		fmt.Println("This is my function")
	}
	myFunc()

	// Here we define another one, but using explicit initialization
	var myFuncExp func() = func() {
		fmt.Println("This is my explicit function")
	}
	myFuncExp()

	// Here we do so again, but we also specify params and returns
	// Definition is just the hypothetical signature
	var divideFloatsVar func(float64, float64) (float64, error)
	// Initialization puts names to params and defines logic
	divideFloatsVar = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	div, err := divideFloatsVar(5.0, 3.0)
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
	fmt.Println("Result", div)
	fmt.Println("")

	/*
		Methods

		Methods are functions that are associated with type
		instances.  We can call them using the dot operator on
		the type for which the method is defined.
	*/
	fmt.Println("#### Methods ####")

	// Initialize a greeter struct instance
	myGreeter := greeter{
		name:     "Dave",
		greeting: "Hello",
	}

	// Call the greet function (copies)
	myGreeter.greet()

	// Call the greetRef function (pointer reference)
	myGreeter.greetRef()
}

/*
Basic syntax

We don't need to go much further than writing a
basic hello world function in order to explore the
basic syntax surrounding go functions.

Functions are identified by the func keyword.

The function name follows, and the casing of the
function name drives visibility in the same way as
is done for variables.

The function name is followed immediately by a pair
of parenthesis.  In this case, the parenthesis are
empty, but in the future we will explore cases in
which parameters are defined within the parenthesis.

Then, the function body is defined within the curly
braces following the function name and parenthesis.
The go compiler enforces that the opening curly brace
is included on the same line as the function signature.
*/
func helloWorld() {
	fmt.Println("Hello, world!")
}

/*
Parameters

Here we define a function which accepts a parameter.
In this case, the function accepts a string which we
define as representing a name, and the function says
hello to that name.
*/
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

/*
Parameters

Here we define a function which accepts multiple parameters.
In this case, the function accepts a string which we define
as representing a name, and an integer which represents the
number of times to say hello to that name.  The function
constructs a loop which says hello to that name n times.
*/
func sayHelloMany(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("Hello, %s!\n", name)
	}
}

/*
Parameters

Here we define a function which accepts multiple like-typed
parameters.  We do so to demonstrate that the go compiler
can infer function parameter types if given as a comma-sep
list followed by the type.
*/
func sayGreeting(greeting, name string) {
	fmt.Printf("%s, %s!\n", greeting, name)
}

/*
Parameters

Here we define a function which mutates one of its parameters.
We show that behind the scenes, the go runtime copies its input
parameters when they are supplied into the function.  So mutating
a parameter does not have any impact beyond the scope of the
function.
*/
func sayAndMutateGreeting(greeting, name string) {
	fmt.Printf("%s, %s!\n", greeting, name)
	greeting = "Goodbye"
}

/*
Parameters

Here we define a function which accpets a pointer which references
some underlying data.  We show that this allows us to mutate the
underlying data inside the function.
*/
func sayAndMutateGreetingRef(greeting *string, name string) {
	fmt.Printf("%s, %s!\n", *greeting, name)
	*greeting = "Goodbye"
}

/*
Parameters

Here we define a function which accepts a variatic int parameter
and sums the results.  We can only define a single variatic
parameter, and it must be the last function parameter.
*/
func sumMany(values ...int) {
	fmt.Println(values) // It is a slice

	// Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

	// For now just print the result
	fmt.Println("The sum is", result)
}

/*
Return values

Here we define a similar function as above, but instead of
printing the result, we return it.  Note that we have to
define the return type in the function signature between the
parenthesis and the curly braces.
*/
func sumManyAndReturn(values ...int) int {
	// Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

	// Return the result
	return result
}

/*
Return values

Here we define another similar function as above, but instead of
returning the result as a copied variable, we can return it as a
pointer.  Under the hood, the go runtime promotes the value to a
common location in memory, and out of the local function stack.
*/
func sumManyAndReturnRef(values ...int) *int {
	// Loop through the values and sum into result
	result := 0
	for _, v := range values {
		result += v
	}

	// Return a pointer to the result
	return &result
}

/*
Return values

Here we define another similar function as above, but we define
the return value as a named parameter.  This initializes the
return value as a zeroed value within the function's scope.
The function is expected to then populate the value.  Then, the
value is implicitly returned at the end of the function.  For
this, we need to be careful of the zero value of the return type.
*/
func sumManyAndReturnImplicitly(values ...int) (result int) {
	// Loop through the values and sum into result
	for _, v := range values {
		result += v
	}
	return // Needs to still be given explicitly
}

/*
Return values

Here we define a division function which divides two floats.  We
know that division by zero is an exceptional case, but how should
we handle this?  Here we show how to write a function which returns
two values, one is the actual return value, and another is an error
value which represents whether an error occurred in that function.
*/
func divideFloats(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
