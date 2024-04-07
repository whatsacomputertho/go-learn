# Looping

Here we explore the `for` keyword in go, and cover the different looping behaviors which can be achieved in go using this keyword.

- [Looping](#looping)
  - [For loops](#for-loops)

## For loops

Here we cover how one may write simple loops using the `for` keyword, how one may exit early from a loop, and how one may loop through a collection type.

We firstly note that the `for` keyword was chosen as the only looping keyword in go.  It handles traditional `while` and `do while` looping behavior as special cases.

For simple `for` loops, we have three basic forms.

**For + initializer + test + incrementer** - Used when we want initialization, incrementation, and break logic handled for us
```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

**For + test** - Used when initialization and incrementation is handled explicitly
```go
i := 0 // i initialized outside loop
for i < 3 {
    fmt.Println(i)
    i++ // Incrementation explicitly defined inside loop
}
```

**For-only** - Used when initialization, iteration, and break logic is handled explicitly
```go
i := 0 // i initialized outside loop
for {
    fmt.Println(i)
    if !(i < 3) {
        break // Break explicitly defined inside loop
    }
    i++ // Incrementation explicitly defined inside loop
}
```

We also covered three concepts which allow us to exit loops and iterations of loops early.  We have

**Break keyword** - Breaks out of a loop entirely
```go
i := 0
for {
    fmt.Println(i)
    if i%2 == 0 {
        break // Break the loop as soon as i is even
    }
}
```

**Continue keyword** - Breaks out of a single iteration of aloop
```go
for i:= 0; i < 6; i++ {
    if i%2 == 0 {
        continue // Skip below logic when i is even
    }
    fmt.Println(i)
}
```

**Labels** - Used to explicitly break out of parent loop from inner loop
```go
MyLoop:
for i :=  0; i < 10; i++ {
    for j := 0; j < 10; j++ {
        if i*j < 3 {
            break MyLoop // Break all loops as soon as i*j < 3
        }
    }
}
```

We finally covered looping over collections using the `range` keyword against a collection in a `for` loop.

**For-range statement**
```go
/*
This will print:
0 1
1 2
2 3
*/
s := []int{1, 2, 3}
for k, v := range s {
    fmt.Println(k, v)
    
}
```