# Constants

This application explores the usage of constants in go.

- [Constants](#constants)
  - [Properties of constants](#properties-of-constants)
  - [Naming conventions for constants](#naming-conventions-for-constants)
  - [Typed and untyped constants](#typed-and-untyped-constants)
  - [Enumerated constants](#enumerated-constants)
  - [Enumerated expressions](#enumerated-expressions)

## Properties of constants

Constants are immutable, but can be shadowed from an inner scope.  Constants must be replaced by the compiler at compile time, and thus the value must be calculable at compile time.  Constants cannot be set to the return value of a function call, for example.

## Naming conventions for constants

Constants are named like variables, not using `UPPER_CASE` casing.  We use
- `PascalCase` for exported constants
- `camelCase` for local constants

## Typed and untyped constants

Typed constants in go work like immutable variables.  They can interoperate only with variables of the same type.

Untyped constants work in exactly the same way as their literal values.  Thus, untyped constants can interoperate with variables of different, but similar types (i.e. int + int16 is okay).

## Enumerated constants

Special symbol `iota` allows for the creation of enumerated constants in which the first in a constant block is an int assigned value `0`, the second is an int assigned value `1`, and so on.  Watch out for constant values that match zero values for variables.

## Enumerated expressions

When defining enumerated constants, we can perform arithmetic and bitwise operations to define the enumeration is initialized, and how it counts.  For example, we can create exponential constants to count file sizes, or perhaps we can create byte-flag constants to define boolean properties which can be combined together in different ways.