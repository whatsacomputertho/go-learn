package main

import (
	"fmt"
)

func main() {
	/*
		Creation of arrays

		When declaring an array, we need to specify both
		the type of variables stored by the array, as well
		as the number of elements stored within the array.

		We can initialize an array after declaring it by
		listing our elements as a comma separated list in
		curly braces following the declaration statement.

		Grouping elements into an array makes accessing the
		like-elements faster as we know they are stored
		contiguously in memory by the go runtime.

		We can use the ... syntax to define the size of the
		array based on the number of elements in the array
		literal, or we can explicitly declare an array to
		have n elements in our declaration statement.

		We can also obviously declare an empty array using
		the declaration statement on one line, and then
		populate it on the following lines.  We do so by
		assigning a variable of the accepted type stored by
		the array to the nth array index.

		We can then access the array's data by index, and
		we can also calculate the length of an array using
		the built in len() function.  This gives the alloc'd
		size of the array regardless of the number of elems.
	*/
	fmt.Println("#### Creation of arrays ####")

	// Array literals with inherited size from literal
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	// Array declaration with subsequent value initialization
	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	students[2] = "Maggie"
	students[1] = "Bart"
	fmt.Printf("Student #1: %v\n", students[1])

	// Built-in len function to get array length
	fmt.Printf("Number of students: %v\n", len(students))
	fmt.Println("")

	/*
		Properties of arrays

		So far we've aggregated primitives under arrays,
		but we can aggregate arbitrary types under arrays.
		Here, we see we can initialize an array of arrays,
		and form a basic fixed-size matrix.

		In many languages, reassignment of an array simply
		results in the establishment of a new pointer to the
		array's data.  In go, reassignment of an array causes
		the full array to be copied into the new variable.

		We need to explicitly point to an array in go if we
		do not wish to copy the array upon reassignment.  In
		this case, mutating the reassigned value results in
		the original value also being mutated, as they point
		to the same underlying data on the heap.

		As we have seen, arrays must have a fixed, known size
		at compile time, which limits their usefulness but
		also is a part of what makes them so efficient.
	*/
	fmt.Println("#### Properties of arrays ####")

	// Arrays can aggregate more than just primitives
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity matrix: %v\n", identityMatrix)

	// Array reassignment & copy behavior
	arr1 := [...]int{1, 2, 3}
	arr2 := arr1
	arr2[1] = 5 // Modifying array 2 doesn't affect array 1
	fmt.Printf("Array 1: %v\n", arr1)
	fmt.Printf("Array 2: %v\n", arr2)

	// Array reassignment using pointers
	arr3 := &arr1
	arr3[2] = 6
	fmt.Printf("Array 1: %v\n", arr1)
	fmt.Printf("Array 3: %v\n", arr3)
	fmt.Println("")

	/*
		Creation of slices

		Slices are closely related to arrays.  We see that
		in the syntax of their creation - we create a slice
		using very similar syntax as array creation, but we
		leave the size of the slize empty rather than stating
		it explicitly or having it inherit from its literal
		definition.

		With the exception of a few things, slices are nearly
		identical to arrays.  We can reuse the built-in len
		function to determine the length of a slice, for
		example.

		However, we find that there is an additional built-in
		function for slices called its capacity.  This can
		differ from the actual observed length of the slice.
		We will later explore the benefits around this feature.

		We can initialize slices as, well, slices of a larger
		array or slice.  We explore the various ways of doing
		this below.

		There is a built-in function called make() in go which
		accepts either 2 or three arguments.  We can use the
		make() function to declare and initialize slices.  The
		benefit of using the make() function is that we can
		explicitly set the slice capacity on initialization.
	*/
	fmt.Println("#### Creation of slices ####")

	// Declaration and initialization of a slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Slice: %v\n", slice)

	// Built-in length and capacity functions
	fmt.Printf("Length of slice: %v\n", len(slice))
	fmt.Printf("Capacity of slice: %v\n", cap(slice))

	// Declaration and initialization of slices as subsets
	// of a larger slice or array
	subSlice1 := slice[:]   // Slice of all elements
	subSlice2 := slice[3:]  // Slice from 4th element to end
	subSlice3 := slice[:6]  // Slice first 6 elements
	subSlice4 := slice[3:6] // Slice the 4th through 6th elements
	fmt.Printf("Sub-slice 1 ([:])  : %v\n", subSlice1)
	fmt.Printf("Sub-slice 2 ([3:]) : %v\n", subSlice2)
	fmt.Printf("Sub-slice 3 ([:6]) : %v\n", subSlice3)
	fmt.Printf("Sub-slice 4 ([3:6]): %v\n", subSlice4)

	// We can also slice an array
	array := [...]int{9, 8, 7, 6}
	arrslice := array[1:3]
	fmt.Printf("Array slice: %v\n", arrslice)

	// Initialization of a slice using make function
	makeSlice := make([]int, 3, 10)
	fmt.Printf("Make slice: %v\n", makeSlice)
	fmt.Printf("Make slice length: %v\n", len(makeSlice))
	fmt.Printf("Make slice capacity: %v\n", cap(makeSlice))
	fmt.Println("")

	/*
		Properties of slices

		Slices are what are known as reference types in go.
		That means that when reassigning slices, we do not
		copy the underlying data in the slice, but instead
		refer back to the underlying data.  So when we
		reassign and mutate a slice, we mutate all references
		to that slice.

		Slices are still fixed-size collections, but we can
		add and remove elements from them.  This makes them
		useful as they are fixed-size but still dynamic.

		If we exceed the capacity of a slice, then the elements
		are copied into a new underlying array of a larger size.
		This becomes expensive as it scales.

		The append function allows us to append an arbitrary
		number of elements into a slice.  It is a variatic
		function.

		The spread operator can be used along with the append
		function in go to result in something equivalent to what
		might be called "extend" in other languages.  It allows
		us to append all elements of another slice into a slice
		in order.

		To treat a slice like a stack, we might want the ability
		to push elements onto the slice, and pop elements off of
		the slice.  We can do so via slicing operations and
		append calls, however we should remain conscious of how
		this impacts the underlying array.
	*/
	fmt.Println("#### Properties of slices ####")

	// Reassignment of slices & copy behavior
	sli1 := []int{4, 5, 6}
	fmt.Printf("Slice 1: %v\n", sli1)
	sli2 := sli1
	sli2[1] = 2
	fmt.Printf("Slice 1: %v\n", sli1)
	fmt.Printf("Slice 2: %v\n", sli2)

	// Appending elements to a slice
	dynamicSlice := []int{}
	fmt.Printf("Dynamic slice: %v\n", dynamicSlice)
	fmt.Printf("Dynamic slice length: %v\n", len(dynamicSlice))
	fmt.Printf("Dynamic slice capacity: %v\n", cap(dynamicSlice))
	dynamicSlice = append(dynamicSlice, 1, 2, 3, 4, 5)
	fmt.Printf("Dynamic slice: %v\n", dynamicSlice)
	fmt.Printf("Dynamic slice length: %v\n", len(dynamicSlice))
	fmt.Printf("Dynamic slice capacity: %v\n", cap(dynamicSlice))

	// Extending a slice using the spread operator
	extraSlice := []int{6, 7, 8}
	dynamicSlice = append(dynamicSlice, extraSlice...)
	fmt.Printf("Dynamic slice: %v\n", dynamicSlice)
	fmt.Printf("Dynamic slice length: %v\n", len(dynamicSlice))
	fmt.Printf("Dynamic slice capacity: %v\n", cap(dynamicSlice))

	// Treating a slice like a stack
	stackSlice := []int{0, 1, 2, 3, 4}
	fmt.Printf("Stack slice: %v\n", stackSlice)
	fmt.Printf("Stack slice length: %v\n", len(stackSlice))
	fmt.Printf("Stack slice capacity: %v\n", cap(stackSlice))
	stackSlice = append(stackSlice, 5) // Push 5 onto stack
	fmt.Printf("Stack slice: %v\n", stackSlice)
	fmt.Printf("Stack slice length: %v\n", len(stackSlice))
	fmt.Printf("Stack slice capacity: %v\n", cap(stackSlice))
	stackSlice = stackSlice[:len(stackSlice)-1] // Pop 5 off of stack
	fmt.Printf("Stack slice: %v\n", stackSlice)
	fmt.Printf("Stack slice length: %v\n", len(stackSlice))
	fmt.Printf("Stack slice capacity: %v\n", cap(stackSlice))
}
