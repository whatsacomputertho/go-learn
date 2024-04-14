# Pointers

Here we explore the usage of pointers in go at a high level.  For now, we do not venture too deeply into the practical usage of pointers as we will explore that in later modules.

- [Pointers](#pointers)
  - [Creating pointers](#creating-pointers)
  - [Dereferencing pointers](#dereferencing-pointers)
  - [Types with internal pointers](#types-with-internal-pointers)

## Creating pointers

In go, using an asterisk just before a type name denotes that the variable will contain a memory address which corresponds to a variable of that type.
```go
// This is not an int, but it is a memory address
// which points to an int in memory
var x *int
```

We can also use the address-of operator `&` to obtain a memory address of an existing variable.
```go
x := 10
fmt.Println(&x) // 0x123abcdef
```

We can do this directly in the initializer as well
```go
// Example struct
type myStruct struct {
	foo int
}

// Main function
func main() {
    // Example of initializing a struct pointer in struct initializer
    ms := &myStruct{foo: 42} // Initialize struct and return pointer
    fmt.Println(ms)          // &{42}
}
```

We can also use the `new()` function to allocate a zeroed instance of a type and return a pointer to that allocated type instance.
```go
// Example struct
type myStruct struct {
	foo int
}

// Main function
func main() {
    // Example of initializing a struct pointer in struct initializer
    ms := new(myStruct) // Initialize struct and return pointer
    fmt.Println(ms)     // &{0}
}
```

## Dereferencing pointers

We can use the asterisk operator in front of a pointer variable to fetch the underlying value to which the pointer refers.
```go
var x *int      // Declare an int pointer
*x = 14         // Initialize the int value using dereferencing
fmt.Println(x)  // 0x123abcdef
fmt.Println(*x) // 14
```

For structs, the compiler is capable of implicitly dereferencing struct pointers so we don't have to clutter our syntax.
```go
// Example struct
type myStruct struct {
	foo int
}

// Main function
func main() {
    // Example of implicitly dereferencing struct pointers
    var ms *myStruct    // Declare a struct pointer
    ms = new(myStruct)  // Allocate a zeroed struct instance
    ms.foo = 42         // Mutate the foo struct property
    fmt.Println(ms.foo) // 42
}
```

## Types with internal pointers

All assignment operations in go are copy operations.  Some types, however, contain pointers under the hood.  Examples include maps and slices; we call these reference types.  When reassigned, their values are still copied, but since they are pointers they still refer to the same underlying data.
```go
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
```