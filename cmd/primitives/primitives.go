package main

import (
	"fmt"
)

func main() {
	/*
		Boolean type

		Most common use of the boolean type in go is around
		logical testing.  Booleans are generated as a result
		of logical equivalency testing like what is seen below.
		We can also explicitly initialize a boolean ourselves
		like what is seen below.

		It's also worth noting that in go, variables are
		initialized to the zero value.  Rather than this being
		some uninitialized memory, for booleans this zero value
		is false, as seen below.
	*/
	fmt.Println("#### Boolean types ####")
	var b bool = true //Explicitly initialized boolean
	t := 1 == 1       //Logical equivalency test initialization
	f := 1 == 2       //Logical equivalency test initialization
	var z bool        //Zero value initialization
	fmt.Printf("b: %v (%T)\n", b, b)
	fmt.Printf("t: %v (%T)\n", t, t)
	fmt.Printf("f: %v (%T)\n", f, f)
	fmt.Printf("z: %v (%T)\n\n", z, z)

	/*
		Integer types

		There are various integer types that we will explore
		below.  An unspecified size int will be determined by
		the platform.  This is the default behavior for ints.

		We have both signed and unsigned integers.  Ints by
		default are signed, but we can explicitly initialize
		integers to be unsigned.
		- int8/uint8/byte
		- int16/uint16
		- int32/uint32
		- int64

		We have the standard arithmetic operations for integers,
		those being addition, subtraction, multiplication,
		division, and the modulus operator.  We cannot receive
		a new type by performing arithmetic on types, and we
		cannot perform arithmetic on two different types.  We
		must do type conversion explicitly.

		We have the standard bitwise operations for integers,
		those being AND, OR, XOR, NAND.  We also have the bit
		shift operations for integers, << and >>.
	*/
	fmt.Println("#### Integer types ####")
	d := 3              //Int size determined by platform, at least 32 bit
	var y byte = 255    //uint8/byte type example
	var u uint16 = 1024 //uint16 type example
	var i int32 = 2048  //int32 type example
	fmt.Printf("d: %v (%T)\n", d, d)
	fmt.Printf("y: %v (%T)\n", y, y)
	fmt.Printf("u: %v (%T)\n", u, u)
	fmt.Printf("i: %v (%T)\n", i, i)
	fmt.Printf("Add d + 10 = %v\n", d+10)
	fmt.Printf("Sub d - 10 = %v\n", d-10)
	fmt.Printf("Mul d * 10 = %v\n", d*10)
	fmt.Printf("Div d / 10 = %v\n", d/10)
	fmt.Printf("Mod d mod 10 = %v\n", d%10)
	fmt.Printf("AND d & 10 = %v\n", d&10)
	fmt.Printf("OR  d | 10 = %v\n", d|10)
	fmt.Printf("XOR d ^ 10 = %v\n", d^10)
	fmt.Printf("NAND d &^ 10 = %v\n", d&^10)
	fmt.Printf("SHF d << 10 = %v\n", d<<10)
	fmt.Printf("SHF d >> 10 = %v\n\n", d>>10)

	/*
		Floating point numbers

		Floating point numbers provide relatively precise
		estimates for the real numbers.  By default, floats
		are initialized as float64s, but may also be float32s.

		We have the standard arithmetic operations for real
		numbers, those being addition, subtraction, multiplication,
		and division.  There is no modulus operator, no bitwise
		operators, and no bit shift operators.
	*/
	fmt.Println("#### Float types ####")
	w := 10.2
	x := 3.7
	fmt.Printf("Add w + x = %v\n", w+x)
	fmt.Printf("Sub w - x = %v\n", w-x)
	fmt.Printf("Mul w * x = %v\n", w*x)
	fmt.Printf("Div w / x = %v\n\n", w/x)

	/*
		Complex numbers

		Complex numbers are an extension of the real numbers
		and are considered primitives in go.  In other languages
		we see complex numbers offloaded to math libraries, but
		not in go.

		We have the standard arithmetic operations for real
		numbers, those being addition, subtraction, multiplication,
		and division.  There is no modulus operator, no bitwise
		operators, and no bit shift operators.

		We also have built in functions for accessing only the
		real part of a complex number, or vice versa for the
		imaginary part of a complex number.  We can also convert
		two floats into a complex number.
	*/
	fmt.Println("#### Complex types ####")
	var c complex64 = 1 + 2i
	var k complex64 = 2 + 5.2i
	fmt.Printf("c: %v (%T)\n", c, c)
	fmt.Printf("Add c + k = %v\n", c+k)
	fmt.Printf("Sub c - k = %v\n", c-k)
	fmt.Printf("Mul c * k = %v\n", c*k)
	fmt.Printf("Div c / k = %v\n", c/k)
	fmt.Printf("real(k): %v (%T)\n", real(k), real(k)) //k is complex64, real part is float32
	fmt.Printf("imag(k): %v (%T)\n", imag(k), imag(k)) //k is complex64, imag part is float32
	//If k was complex128, parts would be float64
	fmt.Printf("complex(w, x): %v (%T)\n\n", complex(w, x), complex(w, x))

	/*
		Text types

		A string encodes UTF-8 characters.

		As we can see by the below, strings are just aliases for
		byte arrays.  When we index a string, we get a byte back,
		which is the uint8 UTF-8 character code of the character.
		We can re-encode the string and print it to visualize the
		character as text.

		A string is immutable in go.  We cannot modify a string
		value once initialized.

		We have a pseudo-arithmetic operation that can be run on
		strings, which is the concatenation operator, +.  We can
		concatenate strings this way.

		We can convert a string directly to a byte array and see
		its UTF-8 character codes represented in array form.  This
		is useful as many common functions operate on byte slices.
		Both the functions for serving a file and serving a response
		operate on byte slices.

		For UTF-32 encoding we have the rune type alias, which is
		the same as an int32.
	*/
	fmt.Println("#### Text types ####")
	s := "this is a string"
	r := "this is also a string"
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("s[2]: %v (%T)\n", s[2], s[2])
	fmt.Printf("string(s[2]): %v (%T)\n", string(s[2]), string(s[2]))
	fmt.Printf("s + r: %v (%T)\n", s+r, s+r)
	by := []byte(s)
	fmt.Printf("[]byte(s): %v (%T)\n", by, by)
	ru := 'a' //A rune is an int32
	fmt.Printf("ru: %v (%T)\n", ru, ru)
}
