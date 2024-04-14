package main

import (
	"fmt"
)

// Used in pointer creation examples
type myStruct struct {
	foo int
}

func main() {
	/*
		Intro to pointers & value types

		Value types in go are types which are copied when
		reassigned to new variables.  We exemplify this with
		the int type.

		We can declare a pointer to establish a duplicate
		reference to the same underlying value type if we wish.
		We must do so explicitly.  Pointer variables are just
		memory addresses which point to a value of that type
		on the heap.  They are declared using the address-of
		operator "&".

		To get the underlying value from the heap, we can
		dereference a pointer variable using the dereferencing
		operator "*".  We can also use this to mutate referenced
		values on the heap.
	*/
	fmt.Println("#### Intro to pointers & value types ####")

	// Basic example of value types in go
	a := 42           // Declare and initialize a
	b := a            // Copy a into b
	a = 27            // Mutate a, this does not affect b
	fmt.Println(a, b) // 27 42

	// Example of pointing to value types
	var c int = 42    // Declare and initialize c
	var d *int = &c   // Point at c using the address-of operator (&)
	c = 27            // Mutate c, this affects d
	fmt.Println(c, d) // 27 0x123abcdef

	// Example of dereferencing a pointer variable
	fmt.Println(*d) // 27

	// Example of mutating heap values using dereferencing
	*d = 14            // Mutate d, this also affects c
	fmt.Println(c, *d) // 14 14
	fmt.Println("")

	/*
		Pointer arithmetic in go

		Pointer arithmetic is a common feature of other
		languages which support pointers as variables.  In
		go, however, we see that pointer arithmetic is not
		supported.

		In advanced scenarios, we can use the built in
		unsafe package in go to deal with more advanced memory
		management.
	*/
	fmt.Println("#### Pointer arithmetic in go ####")

	// Example of pointer arithmetic being unsupported in go
	myArr := [3]int{1, 2, 3}
	arrPtrA := &myArr[0]
	// Trying to hop to arrPtrA, but yields an exception
	// arrPtrB := &myArr[1] - 4 // 4 is sizeof(int)
	fmt.Printf("%v %p\n", myArr, arrPtrA)
	fmt.Println("")

	/*
		Creating pointer types

		Pointer types can be declared explicitly and
		initialized.  We can optionally use the new()
		function to allocate memory for an instance of
		a type and return a pointer to the type. This
		will yield a zeroed value of the type of the
		pointer variable.

		Pointers which are not initialized will contain
		the special value nil.  Nil pointers should be
		handled carefully, as if we try to drill into a
		nil pointer, we will get a runtime exception and
		our program will crash.
	*/
	fmt.Println("#### Creating pointer types ####")

	// Example of creating a struct pointer type
	var ms1 *myStruct        // Declare a struct pointer
	ms1 = &myStruct{foo: 42} // Initialize the struct value
	fmt.Println(ms1)         // &{42}

	// Example of using new to create a struct pointer
	var ms2 *myStruct   // Declare a struct pointer
	ms2 = new(myStruct) // Allocate a zeroed struct instance
	fmt.Println(ms2)    // &{0}

	// Example of creating a nil pointer
	var ms3 *myStruct // Declare a struct pointer, don't init
	fmt.Println(ms3)  // <nil>
	fmt.Println("")

	/*
		Dereferencing pointer types

		In order to get at the underlying data corresponding
		to a pointer type, we will need to do what is called
		dereferencing.  This converts a pointer to an instance
		of its corresponding type by fetching the underlying
		value from the heap.

		It turns out that we don't need to explicitly deref
		structs when referencing them using a pointer.  The
		compiler allows us to access struct properties through
		a struct pointer implicitly.  It handles the deref
		under the hood.
	*/
	fmt.Println("#### Dereferencing pointer types ####")

	// Example of explicitly dereferencing struct pointers
	var ms4 *myStruct   // Declare a struct pointer
	ms4 = new(myStruct) // Allocate a zeroed struct instance
	(*ms4).foo = 27     // Mutate the foo struct property
	fmt.Println((*ms4).foo)

	// Example of implicitly dereferencing struct pointers
	var ms5 *myStruct   // Declare a struct pointer
	ms5 = new(myStruct) // Allocate a zeroed struct instance
	ms5.foo = 42        // Mutate the foo struct property
	fmt.Println(ms5.foo)
	fmt.Println("")

	/*
		Reference types in go

		We have seen numerous examples of value types versus
		reference types in go.  Here we use arrays and slices
		to exemplify the notion of this distinction.

		We note that slices and maps in particular are reference
		types, and we should be careful when passing reference
		types around across our application as they may mutate
		in unexpected ways.

		Primitives, arrays, and structs on the other hand are
		value types and will be copied when reassigned.  These
		are safer to pass around your application, but note the
		copy may be inefficient.
	*/
	fmt.Println("#### Reference types ####")

	// Example of arrays as value types in go
	myValArr := [3]int{1, 2, 3}        // Initialize array
	otherValArr := myValArr            // Initialize another array from previous array
	fmt.Println(myValArr, otherValArr) // [1, 2, 3] [1, 2, 3]
	myValArr[0] = 42                   // Mutate original array
	fmt.Println(myValArr, otherValArr) // [42, 2, 3] [1, 2, 3]

	// Example of slices as reference types in go
	myRefSlc := []int{1, 2, 3}         // Initialize slice
	otherRefSlc := myRefSlc            // Initialize another slice from previous slice
	fmt.Println(myRefSlc, otherRefSlc) // [1, 2, 3] [1, 2, 3]
	myRefSlc[0] = 42                   // Mutate original slice
	fmt.Println(myRefSlc, otherRefSlc) // [42, 2, 3] [42, 2, 3]

	// Example of maps as reference types in go
	myRefMap := map[string]string{
		"foo":  "bar",
		"fizz": "buzz",
	}
	otherMap := myRefMap               // Initialize another map from previous map
	fmt.Println(myRefMap, otherMap)    // map[foo:bar fizz:buzz] map[foo:bar fizz:buzz]
	myRefMap["foo"] = "qux"            // Mutate original map
	fmt.Println(myRefSlc, otherRefSlc) // map[foo:qux fizz:buzz] map[foo:qux fizz:buzz]
}
