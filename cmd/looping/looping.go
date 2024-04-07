package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Basic for loops

		Here we cover the most basic form of a for loop.
		We have a loop that iterates 5 times, and prints
		the iteration number on each iteration.

		The first statement initializes the counter variable.
		The second statement defines a condition for which
		the loop breaks.  The third statement defines how the
		counter variable is incremented after each iteration.

		We can define multiple counter variables, use them in
		a break condition, and increment them in a single for
		loop.

		We can initialize a variable outside the loop and
		leave the initializer blank in the for loop.  However
		we must be sure to add an initial semicolon to denote
		that our initializer is blank.  We also note that in
		this case, the counter is scoped to the parent scope
		(main function here) versus in the prior cases, the
		counter is only scoped to the individual for loop.
	*/
	fmt.Println("#### Basic for loops ####")

	// Basic go for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Iteration " + strconv.Itoa(i))
	}

	// For loop with multiple counter variables
	for i, j := 0, 0; i < 3 && j < 3; i, j = i+1, j+1 {
		fmt.Println("i, j: " + strconv.Itoa(i) + ", " + strconv.Itoa(j))
	}

	// For loop with counter defined outside the initializer
	ctr := 0
	for ; ctr < 3; ctr++ {
		fmt.Println("ctr: " + strconv.Itoa(ctr))
	}
	fmt.Println("")

	/*
		For loops as while loops

		Similarly to the above, the iterator is optional.
		It can be left out for an infinite loop (equivalent
		to while), or it can be supplied in the for loop
		block explicitly.

		If we leave out the initializer and the iterator,
		this is go's equivalent of a while loop.  We can
		supply only a break condition with no semicolons.

		We can also go further to leave out the break
		condition itself and just initialize the for loop
		blankly.  In this case, we guarantee that we will
		receive an infinite loop.  We must explicitly
		provide break conditions and such within the loop
		block.
	*/
	fmt.Println("#### For loops as while loops ####")

	// For loop with initializer and blank iterator
	for i := 0; i < 3; {
		fmt.Println("Iteration ", strconv.Itoa(i))
		i++ // Iterator given explicitly in loop
	}

	// Infinite loop with initializer, break condition, blank iterator
	//for i := 0; i < 3; {
	//	  fmt.Println("Iteration ", strconv.Itoa(i))
	//}

	// For loop with only break condition, semicolons given explicitly
	//ctr = 0
	//for ;ctr < 3; {
	//	fmt.Println("ctr: " + strconv.Itoa(ctr))
	//	ctr++ // Iterator given explicitly in loop
	//}

	// Equivalent to the above, semicolons can be left out
	ctr = 0
	for ctr < 3 {
		fmt.Println("ctr: " + strconv.Itoa(ctr))
		ctr++ // Iterator given explicitly in loop
	}

	/*
		Break and continue

		Here we explore the break and continue keywords.
		These keywords are used to skip individual iter-
		ations, and to exit a loop altogether.

		We notice that the break keyword only exits the
		innermost loop when used in a nested loop.  We can
		use labels to identify which loop to break when
		breaking from an inner loop.
	*/

	// Infinite loop with explicit break condition
	ctr = 0
	for {
		fmt.Println("ctr: " + strconv.Itoa(ctr))
		if ctr%2 == 0 {
			fmt.Println("ctr is even, breaking")
			break
		}
		ctr++
	}

	// Loop using continue to skip even numbers
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Iteration is odd: " + strconv.Itoa(i))
	}

	// Nested loop using break to exit inner loop
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i, j: " + strconv.Itoa(i) + ", " + strconv.Itoa(j))
			if j%2 == 0 {
				break // Only breaks inner loop
			}
		}
	}

	// Nested loop using break & labels to exit outer loop
MyLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i, j: " + strconv.Itoa(i) + ", " + strconv.Itoa(j))
			if j%2 == 0 {
				break MyLoop // Only breaks inner loop
			}
		}
	}
	fmt.Println("")

	/*
		Looping and collection types

		Here we show how we can loop through collection types
		using a for-range loop in go.  The range keyword gives
		us the key or index of each element, and the value of
		each element inside the collection.  We exemplify this
		with various collection types.

		However, we note that we get an error when we only
		want either the keys or the values of a collection,
		but we assign both to variables.  For this, we can
		use the write-only operator "_" to assign to the
		value we don't need.

		We finally note that there is special behavior when
		looping through channels (multithreading construct)
		which we will revisit later in the tutorial.
	*/
	fmt.Println("#### Looping and collection types ####")

	// Loop through a slice of integers using for-range loop
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println("Index: " + strconv.Itoa(k) + "; Value: " + strconv.Itoa(v))
	}

	// Loop through a map using for-range loop
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println("Key: " + k + "; Value: " + strconv.Itoa(v))
	}

	// Loop through a string using for-range loop
	myStr := "Hello, go!"
	for k, v := range myStr {
		fmt.Println("Index: " + strconv.Itoa(k) + "; Value: " + string(v))
	}

	// Loop through a collection but only use the keys
	for k, _ := range statePopulations {
		fmt.Println("State: " + k)
	}
}
