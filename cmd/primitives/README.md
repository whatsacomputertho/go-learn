# Primitives

This application explores the various primitive types in go, such as those which follow.
- [Primitives](#primitives)
  - [Boolean type](#boolean-type)
  - [Numeric types](#numeric-types)
    - [Integers](#integers)
    - [Floating point numbers](#floating-point-numbers)
    - [Complex numbers](#complex-numbers)
  - [Text types](#text-types)

## Boolean type
Booleans are true/false values, and they are not type aliases.  Their zero value is false.

## Numeric types
We have integers, floating point numbers, and complex numbers.

### Integers

Integers can be signed and unsigned.  The signed integers range from int8 (8-bit) to int64 (64-bit).  The unsigned integers range from byte/uint8 (8-bit) to uint32 (32-bit).

Between integers of the same types, we can perform the following arithmetic and bitwise operations
- Addition
- Subtraction
- Multiplication
- Division
- Modulus
- AND
- OR
- XOR
- NAND
- Bitwise shift left
- Bitwise shift right

### Floating point numbers

Floating point numbers follow the IEEE-754 standard.  Their zero value is 0.  We have float32 and float64.  Their literal styles include decimal (3.14), exponential (13e18 or 2E10), and mixed (13.7e12).

Between floating point numbers of the same types, we can perform the following arithmetic operations
- Addition
- Subtraction
- Multiplication
- Division

### Complex numbers

Complex numbers are uniquely considered primitives in go.  Their zero value is 0+0i.  We have complex64 and complex128.  We also have built in functions for converting between complex and real numbers
- complex - make 2 like-typed floats complex
- real - get real part as float
- imag - get imaginary part as float

Between complex numbers of the same types, we can perform the following arithmetic operations
- Addition
- Subtraction
- Multiplication
- Division

## Text types

For text we have the string type, which follows a UTF-8 character encoding.  Strings are immutable in go, but can be concatenated using the pseudo-arithmetic operation `+`.  Strings are type aliases for the byte slice/byte array `[]byte`.

We also have the rune type, which follows a UTF-32 character encoding.  Runes alias the int32 type.