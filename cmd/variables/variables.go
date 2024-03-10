package main

import (
	"fmt"
	"strconv"
)

/*
Package-level declaration

Here we cannot use implicit declaration, we must explicitly
specify the type at declaration time.  We can use a var block
to initialize many variables at once, like so.
*/
var (
	title       string = "Variables"
	description string = "In this module, we're learning about go variables"
	n           int    = 44
)

/*
Package-level variable naming conventions

- Pascal or camelCase where acronyms are capitalized

Lowercase package-level variable names denote a package-scoped
variable.  Uppercase package-level variable names denote an
externally-scoped variable.  Block-level variables are never
externally-scoped.  Package-level variables can go unused without
the compiler yelling at you.

In general, the length of a variable name should represent the
lifetime of the variable.  For loops and such, i is okay.  For
long-lived variables, we want to be a bit more verbose for clarity
and uniqueness.  Still try to remain concise in the process.

For acronyms, the best practice in go is to leave acronyms in
uppercase.  A variable named myHttpRequest is not best practice.
It should be named myHTTPRequest.
*/
var myvar int = 1
var MYVAR int = 2

func main() {
	//Print the package-level title and description variables
	fmt.Println(title)
	fmt.Println(description)

	/*
	   Multiline declaration

	   Declare a new integer variable on one line, then initialize
	   its value on the next line.  This is useful for declaring a
	   variable in a parent scope, then initializing its value in
	   one of its child scopes.
	*/
	var i int
	i = 27
	i = 33 //Mutation is possible
	fmt.Printf("i: %v (%T)\n", i, i)

	/*
	   Single-line explicit declaration

	   Declare and initialize a new integer variable on one line.
	   Here, we explicitly state the type of the variable.  This
	   is useful if the literal value needs to be cast into the
	   specified type.
	*/
	var j int = 27
	var f float64 = 27 //Casting int literal to float64
	fmt.Printf("j: %v (%T)\n", j, j)
	fmt.Printf("f: %v (%T)\n", f, f)

	/*
	   Single-line implicit declaration

	   Declare and initialize a new integer variable on one line.
	   Here, we do not explicitly state the type of variable.  This
	   is
	*/
	k := 54
	fmt.Printf("k: %v (%T)\n", k, k)

	/*
		Redeclaration and shadowing

		Variables may not be redeclared, but they may be shadowed.
		Here we see that we can shadow the package-level n variable
		in this child scope.  We cannot use implicit declaration here.
	*/
	fmt.Printf("n: %v (%T)\n", n, n) //Printing package-level n
	var n int = 57                   //Shadowing n from package-level scope
	//n := 37 //Uncomment this, this will cause an error
	fmt.Printf("n: %v (%T)\n", n, n)

	/*
		Unused variables

		If you do not use variables, the go compiler will yell at
		you.  Uncomment the printf statement and see for yourself.
	*/
	u := 22
	fmt.Printf("u: %v (%T)\n", u, u)

	/*
		Type casting

		Here we experiment with type casting in go.  We can convert
		types explicitly by using the conversion function type(var).

		Go allows us to explicitly convert types but it does not
		allow us to implicitly convert types.  This way it is the
		programmer's responsibility to understand when information
		is lost in conversion.

		Casting integers to strings does not quite work as expected.
		A string is an alias for a stream of bytes.  Thus, the conv
		looks for the unicode character at 42 when initializing the
		string.
	*/
	var m float64
	m = float64(i) //Casting i into a float64 and storing in m
	fmt.Printf("m: %v (%T)\n", m, m)
	var o int
	o = int(f) //Casting f into an int explicitly, this works fine
	fmt.Printf("o: %v (%T)\n", o, o)
	//o = f //This doesn't work, uncomment to see for yourself
	var s string
	s = string(o)                    //Casting o into a string explicitly
	fmt.Printf("s: %v (%T)\n", s, s) //Actually prints as unicode, not int val
	var t string
	t = strconv.Itoa(o)
	fmt.Printf("t: %v (%T)\n", t, t) //Prints int value
}
