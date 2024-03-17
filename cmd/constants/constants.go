package main

import (
	"fmt"
	//	"math"
)

const PublicConst int = 27 // Exported

// See enumerated constants
const (
	i = iota
	j = iota
	k // This is ok thanks to compiler inferencing
)

// See enumerated constants
const (
	i2 = iota
)

// See enumerated constants
const (
	_ = iota + 10
	labrador
	corgi
	pointer
	shepard
)

// See enumerated constants
const (
	_  = iota             // Ignore first value
	KB = 1 << (10 * iota) // Bit shift each by 10
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// See enumerated constants
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

// See enumerated constants

func main() {
	/*
		Naming conventions for constants

		Constants in go are generally named the same way as
		variables, using camelCase.  We do not name them in
		the same way as other languages (UPPER_CASE) since
		variables with an uppercase first letter are exported
		publicly.
	*/
	fmt.Println("#### Naming conventions for constants ####")

	// Naming conventions for constants
	const privateConst int = 33 // Not exported
	fmt.Printf("privateConst: %v (%T)\n", privateConst, privateConst)
	fmt.Printf("PublicConst: %v (%T)\n", PublicConst, PublicConst)
	fmt.Println("")

	/*
		Properties of constants

		Mutating a constant leads to a compiler error.  If a
		constant is assigned a value which requires a function
		to be executed in determining that value, a compiler
		error is thrown.  All constants must be defined at
		compile time.

		Constants can store any primitive value.  However,
		constants cannot store collections in go as collections
		are mutable by default (as we'll learn).

		Constants can be shadowed in a child scope.  If a const
		is available in a parent scope, the child scope can
		re-declare that constant with a different value and a
		different type.  This is valid in go.

		Constants can be used alongside variables of the same
		type in things like arithmetic.  The result is a var
		of the appropriate type and not a constant.
	*/
	fmt.Println("#### Properties of constants ####")

	// Mutating constants
	//privateConst = 25 // Compiler error
	//const errorConst float64 = math.Sin(1.57) // Compiler error

	// Constants can store primitives
	const intConst int = 25
	const strConst string = "foo"
	const floConst float64 = 3.14
	const booConst bool = true
	fmt.Printf("intConst: %v (%T)\n", intConst, intConst)
	fmt.Printf("strConst: %v (%T)\n", strConst, strConst)
	fmt.Printf("floConst: %v (%T)\n", floConst, floConst)
	fmt.Printf("booConst: %v (%T)\n", booConst, booConst)

	// Constants can be shadowed
	const PublicConst byte = 12
	fmt.Printf("PublicConst: %v (%T)\n", PublicConst, PublicConst)

	// Constant-variable arithmetic
	var byteVar byte = 10
	fmt.Printf("PublicConst + byteVar = %v\n", PublicConst+byteVar)
	fmt.Println("")

	/*
		Untyped constants

		Constants can be declared using the compiler's type
		inferencing feature in which no explicit type is
		given but a value is assigned.  However, unlike vars,
		constants when untyped are treated as literals, and
		can be more fluid in terms of their typing as seen
		in the below proof-of-concept.
	*/
	fmt.Println("#### Untyped constants ####")

	// Untyped constants
	const a = 42                             // Defaults to int
	fmt.Printf("a: %v (%T)\n", a, a)         // Confirms type
	var b int16 = 27                         // Not int, int16
	fmt.Printf("a + b: %v (%T)\n", a+b, a+b) // But this is okay
	fmt.Println("")

	/*
		Enumerated constants

		Enumerated constants relate to the iota keyword in
		the go programming language.  If we assign a const
		to iota in a constant block (see above) it is set to
		an int with value 0.

		If we assign more consts to iota, then they evaluate
		to 1, then 2, and so on. This is true even when the
		latter constants are not assigned to iota explicitly,
		thanks to the compiler's inferencing features.

		Iota is scoped to a single constant block.  It resets
		when we enter a new block.

		Example usage of iota is to establish an enumerated
		typing of a struct/object.  Say we have a Dog struct
		which can contain a breed.  Maybe we use an enum const
		to define the possible breeds and to assign the Dog
		instances a breed.

		One warning with the above is with zero-values.  Since
		iota starts at zero, if we declare but do not init a
		type var, then it will match the first entry in the
		enum const, which may seem unexpected.

		One way around the above is to use the zero-value as
		an error value in the enum so that these edge cases
		are caught.  Or we can use the underscore character
		for the initial value (go's only write-only variable)
		to ignore it.

		Within constant blocks, we can do a limited amount of
		dynamic things, like arithmetic and bitwise operations
		to define where our enum starts counting from, and how
		it counts.
	*/
	fmt.Println("#### Enumerated constants ####")

	// Enumerated constants
	fmt.Printf("i: %v (%T)\n", i, i)
	fmt.Printf("j: %v (%T)\n", j, j)
	fmt.Printf("k: %v (%T)\n", k, k)

	// Separate constant block with iota
	fmt.Printf("i2: %v (%T)\n", i2, i2)

	// Enumerated constants for types of things
	var breed int = labrador
	fmt.Printf("breed: %v (%T)\n", breed, breed)
	fmt.Printf("is labrador? %v\n", breed == labrador)

	// Bit shifting for exponential enums
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// Bit shifting for boolean flags in a byte
	// Roles stored in a byte, or-ed into user "roles" byte
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("Roles byte: %b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is at hq? %v\n", isHeadquarters&roles == isHeadquarters)
}
