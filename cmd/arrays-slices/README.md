# Arrays and Slices

This application explores the various functionality surrounding arrays and slices in go.

- [Arrays and Slices](#arrays-and-slices)
  - [Arrays](#arrays)
  - [Slices](#slices)

## Arrays

At a high level, we explore how arrays are created, as well as the built-in functions for handling and interacting with arrays.

Arrays are collections of variables of the same type.  Arrays do not need to operate on only primitives.  Arrays are fixed-size; their size must be defined at compile time, and their size cannot change at run time.

We can declare arrays as literals with inherited size, literals with explicit size, or simply without initialization
```go
a := [3]int{1, 2, 3}
b := [...]int{1, 2, 3}
var c [3]int
```

As is standard in most other languages, array data is accessed via a zero-based index.  We pass an offset relative to the first memory address of the array to access the nth element of the array.
```go
a := [3]int{1, 2, 3}
fmt.Println(a[1]) // 2
```

We can use the `len()` function (built-in in go) to calculate the number of elements for which an array is sized.
```go
a := [3]int{1, 2, 3}
fmt.Println(len(1)) // 3
```

Array reassignment results in a copy being made of the underlying array.  Mutations on the array copies do not affect the original array.  This can become very expensive as the array size scales.
```go
a := [3]int{1, 2, 3}
b := a
b[1] = 5
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [1 5 3]
```

We can use pointer references to point at the same underlying array as opposed to copying, but should be aware of the underlying memory representation to avoid unforeseen bugs when making mutations.
```go
a := [3]int{1, 2, 3}
b := &a
b[1] = 5
fmt.Println(a) // [1 5 3]
fmt.Println(b) // &[1 5 3]
```

## Slices

Similarly for slices, we explore how slices are created, as well as the built-in functions for handling and interacting with slices.

Every slice in go is backed by an array under the hood.  Slices can be thought of in numerous ways.  First, we might think of a slice as a subset of an array or parent slice.  Or, we might think of a slice as a dynamic collection but with a fixed-size capacity.

We can initialize a slice as a subset of a parent array/slice.
```go
arr = [...]int{0, 1, 2, 3, 4}
slc := arr[1:3]
fmt.Println(slc) // [1 2]
```

We can initialize a slice as a literal
```go
slc := []int{0, 1, 2, 3}
```

Or we can use the make function to allocate a slice given its type, number of elements, and optionally its capacity.
```go
slc := make([]int, 3)     // Length 3 capacity 3
slc := make([]int, 3, 10) // Length 3 capacity 10
```

The `len()` function can be reused for slices to determine the current length/number of elements in the slice.  There is also a capacity function `cap()` which returns the total number of elements that slice can potentially fit.
```go
slc := make([]int, 3, 10)
fmt.Println(len(slc)) // 3
fmt.Println(cap(slc)) // 10
```

The `append()` function can be used to append new elements to a slice.
```go
slc := make([]int, 3, 10)
slc[0] = 0
slc[1] = 1
slc[2] = 2
fmt.Println(slc) // [0 1 2]
slc = append(slc, 3, 4, 5)
fmt.Println(slc) // [0 1 2 3 4 5]
```

If by appending to a slice we exceed the slice's capacity, then the underlying array of the slice is copied into a larger array that can fit the extra elements.  This can become expensive as the slice size scales.
```go
slc := []int{1, 2, 3}
fmt.Println(cap(slc))      // 3
slc = append(slc, 4, 5, 6) // Triggers a copy operation
fmt.Println(cap(slc))      // Something > 3
```

Lastly, unlike arrays, slices are reference types, and so they refer to the same underlying data when reassigned.
```go
a := []int{1, 2, 3}
b := a
b [1] = 5
fmt.Println(a) // [1 5 3]
fmt.Println(b) // [1 5 3]
```