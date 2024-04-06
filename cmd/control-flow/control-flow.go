package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		If statements

		If statements in go are generally accompanied by
		a boolean input.  If that boolean evaluates to
		true, then some logic inside the statement is
		executed.  Otherwise, it is not executed.

		We use very tangible examples using literals to
		start, but the utility of if statements comes from
		the usage of dynamic statements which evaluate to
		booleans.

		A common idiom in go is the use of initializer
		syntax in if statements.  We use an example of a
		map to exemplify this.
	*/
	fmt.Println("#### If statements ####")

	// Example of an if statement which executes
	if true {
		fmt.Println("Hello, world!") // This will execute
	}

	// Example of an if statement which does not execute
	if false {
		fmt.Println("Goodbye, world!") // This will not execute
	}

	// Example of initializer syntax in an if statement using
	// map key validation
	ages := map[string]int{
		"Joe":  31,
		"Bob":  45,
		"Dave": 59,
	}

	// Initialize the age and ok variables, then pass the ok
	// result in to validate whether the "Joe" key existed
	if joeAge, ok := ages["Joe"]; ok {
		fmt.Printf("joeAge: %v\n", joeAge)
	}
	//fmt.Println(joeAge) // This will fail as joeAge is out of scope

	// Same as above but this time with a nonexistent key
	if robAge, ok := ages["Rob"]; ok {
		fmt.Printf("robAge: %v\n", robAge) // This will not execute
	}
	fmt.Println("")

	/*
		Comparison operators

		We begin by exploring the partial ordering comparison
		operators ">" and "<", as well as the equivalence
		comparison operator "==".  We do so using a naive
		hardcoded number guessing game.

		We also give very basic demonstrations via some printf
		statements to show the usage of the loose partial
		ordering comparison operators ">=" and "<=" as well as
		the inverse equivalence comparison operator "!=".

		We also cover an edge case in which floating point
		arithmetic imprecision is used to express caution
		when comparing floating point numbers.
	*/
	fmt.Println("#### Comparison operators ####")

	// The expected number and our hardcoded "guess"
	number := 50
	guess := 30

	// Comparing our guess against the expected number
	if guess < number {
		fmt.Println("Too low") // If our guess is less than
	}
	if guess > number {
		fmt.Println("Too high") // If our guess is greater than
	}
	if guess == number {
		fmt.Println("Correct") // If our guess is equal to
	}

	// Also note these related comparison operators
	fmt.Printf("Is guess greater than or equal to number? %v\n", guess >= number)
	fmt.Printf("Is guess less than or equal to number? %v\n", guess <= number)
	fmt.Printf("Is guess greater not equal to number? %v\n", guess != number)

	// Example of floating point imprecision
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("They are equal") // This should fire but doesn't
	} else {
		fmt.Println("They are not equal") // This fires due to floating point imprecision
	}
	fmt.Println("")

	/*
		Logical operators

		Here we show how comparison operations can be combined
		logically via logical operators to construct dynamic
		control flow logic.

		Suppose later on we wanted to refactor our number
		guessing game to take CLI input from the user.  Then
		we would want to run validation on that input to ensure
		that invalid values are not given in an effort to break
		our system.

		The OR operator "||" returns true if either of its inputs
		evaluate to true.  The AND operator "&&" returns true if
		all of its inputs evaluate to true.  The NOT operator
		flips its input boolean.

		We also cover short circuiting, in which the program
		enters the if block as soon as the condition is guaranteed
		to evaluate to true - it does not always check all conditions!
	*/
	fmt.Println("#### Logical operators ####")

	// The expected number and our hardcoded "guess"
	anotherNumber := 90
	anotherGuess := -5

	// Running some validation on our guess
	// Combining checks together via the "||" (OR) operator
	if anotherGuess < 1 || anotherGuess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	}

	// Running similar validation on our guess
	// Combining checks together via the "&&" (AND) operator
	// Will not run as-is since guess is out of range
	if anotherGuess >= 1 && anotherGuess <= 100 {
		if anotherGuess < anotherNumber {
			fmt.Println("Too low")
		}
		if anotherGuess > anotherNumber {
			fmt.Println("Too high")
		}
		if anotherGuess == anotherNumber {
			fmt.Println("Correct")
		}
	}

	// Basic demonstration of the "!" (NOT) operator
	fmt.Printf("!true: %v\n", !true) // false

	// Demonstration of short circuiting
	// We might expect returnTrue to run and print "Returning true"
	// But it doesn't due to short circuting, once anotherGuess < 1
	// evaluates to true, we proceed to the if block logic
	if anotherGuess < 1 || returnTrue() || anotherGuess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	}
	fmt.Println("")

	/*
		If / else if / else statements

		Here we explore way in which we can refactor the
		above for cleanliness so that we do not repeat our
		logical test.  We use if / else if / else to do so.
	*/
	fmt.Println("#### If / else if / else statements ####")

	// The expected number and our hardcoded "guess"
	yetAnotherNumber := 20
	yetAnotherGuess := 20

	// Using if / else to run one block if a condition is met,
	// or run another block if that same condition is not met
	if yetAnotherGuess < 1 || yetAnotherGuess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else {
		// Using if / else if / else to run one of many blocks
		// based on mutually exclusive conditions
		if yetAnotherGuess < yetAnotherNumber {
			fmt.Println("Too low")
		} else if yetAnotherGuess > yetAnotherNumber {
			fmt.Println("Too high")
		} else { // Logically that this means guess equals number
			fmt.Println("Correct")
		}
	}
	fmt.Println("")

	/*
		Switch statements

		A switch/case statement is another control flow
		construct in go.  We pass a tag into a switch statement
		and compare it against a pre-defined number of cases.
		We can also define a default case if none of the cases
		are satisfied.

		In go, we have the ability to also define multiple
		comparisons in a single case statement, as seen in the
		below example.  However, we note comparisons must be
		unique, or else we will face a "duplicate case" syntax
		error.

		We can use initializer syntax in a switch statement
		as seen in the second example.  We can also use what
		is known as tagless syntax to run comparisons against
		variables in scope.  We note that overlapping comparison
		is tolerable in tagless switch statements, and the first
		case that is satisfied is executed.

		We also note that break keywords in switch statements
		are implicit in go.  We also note that falling through
		is not implicit in switch statements in go, but we can
		define them explicitly to achieve fallthrough in our
		cases.  However, fallthrough is logicless; the next
		case will execute regardless of if its condition is met
		or not.

		Lastly we cover the type switch.  In go, we can perform
		a switch statement on the type of an interface.  Interfaces
		can be assigned to arbitrary types, and the type of an
		interface can be extracted and used as a tag for a switch
		statement.
	*/
	fmt.Println("#### Switch statements ####")

	// An example switch statement
	switch 2 {
	case 1, 5, 10:
		fmt.Println("One, five, or ten")
	case 2, 4, 6:
		fmt.Println("Two, four, or six")
	//case 5:
	//   fmt.Println("Five") // Duplicate case syntax error
	default:
		fmt.Println("Another number")
	}

	// Initializer syntax in a switch statement
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("One, five, or ten")
	case 2, 4, 6:
		fmt.Println("Two, four, or six")
	default:
		fmt.Println("Another number")
	}

	// Tagless syntax in a switch statement
	// Fallthrough & implicit break keyword
	myVar := 10
	switch {
	case myVar <= 10:
		fmt.Println("myVar is less than or equal to 10")
		fallthrough // Explicit fallthrough - logicless
	case myVar > 20:
		fmt.Println("myVar is greater than 20")
		// Implicit break - no need to explicitly supply break keyword
	case myVar <= 20: // Overlap with first case but still ok
		fmt.Println("myVar is less than or equal to 20")
	default:
		fmt.Println("myVar is greater than 20")
	}

	// Type switch in go
	var myInter interface{} = 1
	switch myInter.(type) {
	case int:
		fmt.Println("myInter is an int")
	case float64:
		fmt.Println("myInter is a float64")
	case string:
		fmt.Println("myInter is a string")
	default:
		fmt.Println("myInter is another type")
	}
}

/*
Logical operators (cont.)

Used in our above example of short circuiting to show
that this function does not run and print "Returning true"
if a prior condition guarantees the larger logically
combined condition to evaluate to true.
*/
func returnTrue() bool {
	fmt.Println("Returning true") // Not printed in our short circuiting example
	return true
}
