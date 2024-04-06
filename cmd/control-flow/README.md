# Control Flow

This application explores the capabilities within go surrounding control flow.  This includes largely the `if` and `switch` statements.

- [Control Flow](#control-flow)
  - [If statements](#if-statements)
  - [Switch statements](#switch-statements)

## If statements

During our investigation into if statements in go, we focus on the various comparison operators which are generally used to drive control flow.  We then cover the `if`, `else if`, and `else` statements.

If statements allow us to conditionally execute a block of code based on a boolean argument.
```go
if true {
    fmt.Println("Hello, world!") // This will run
}

if false {
    fmt.Println("Goodbye, world!") // This will not run
}
```

We can use comparison operators to dynamically execute code blocks based on the values of variables.  These include
- The partial ordering comparison operators `>` and `<`
- The equivalence comparison operator `==`
- The loose partial ordering comparison operators `>=` and `<=`
- The inverse equivalence comparison operator `!=`.
```go
i := 10
if i < 20 {
    fmt.Println("i is less than 20")
}
```

We can make our comparisons even more dynamic by combining multiple comparison operators logically using the logical operators.  These include
- The AND operator `&&`
- The OR operator `||`
- The NOT operator `!`
```go
i := 10
j := 15
if i < 20 && j < 20 {
    fmt.Println("i and j are both less than 20")
}
```

However we should remain aware of short circuiting when combining multiple comparisons.  This is to save operations, as
- OR is satisfied as soon as any of its operands evaluate to `true`
- AND fails as soon as any of its operands evaluate to `false`
```go
i := 10
j := 15
// Here, i < 20 is satisfied
// Thus the entire statement is true
// so j < 20 is ignored
if i < 20 || j < 20 {
    fmt.Println("i and/or j is less than 20")
}
```

We also note that we can fanout if statements into if / else if / else statements for further dynamic behavior when dealing with control flow.
```go
i := 10
if i < 0 {
    fmt.Println("i is negative")
} else if i > 0 {
    fmt.Println("i is positive") // This will execute
} else {
    fmt.Println("i is neither negative nor positive")
}
```

We finally state caution when comparing floats for equality as floats can display imprecision in some cases.
```go
i := 0.123
if i == math.Pow(math.Sqrt(i), 2) {
    fmt.Println("They are equal") // This should fire but doesn't
} else {
    fmt.Println("They are not equal") // This fires due to floating point imprecision
}
```

## Switch statements

During our investigation into switch statements in go, we focus on basic usage of switch/case statements, we then investigate cases with multiple tests, and falling through.  We also cover type switching in go.

We can switch on a tag to perform control flow based on a value.
```go
switch 2 {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two") // This will execute
default:
    fmt.Println("Not one or two")
}
```

We can use initializer syntax in a switch statement in go.  We can also have case statements with multiple tests in go.
```go
switch i := 2 + 3; i {
case 1, 5, 10:
    fmt.Println("One, five, or ten") // This will execute
case 2, 4, 6:
    fmt.Println("Two, four, or six")
default:
    fmt.Println("Some other value")
}
```

We also have tagless switch statements in go.  We can explicitly perform comparison in a switch case statement against an in-scope variable.  We also note that case blocks implicitly break in go, we do not need to provide a `break` keyword ourselves.
```go
i := 10
switch {
case i < 0:
    fmt.Println("i is negative")
case i <= 10: // Overlaps with above case, but this is ok
    fmt.Println("i is less than or equal to ten") // This will execute
    // Implicit break, no further cases will execute
default:
    fmt.Println("i is greater than ten")
}
```

Go supports a logicless fallthrough statement, where if a certain case is met, we can explicitly fall through to the next case.
```go
i := 10
switch {
case i < 0:
    fmt.Println("i is negative")
case i <= 10:
    fmt.Println("i is less than or equal to ten") // This will execute
    fallthrough // Fallthrough to the next case
case i > 20000000: // This case is not satisfied
    fmt.Println("i is really big") // This will execute anyway
default:
    fmt.Println("i is built different")
}
```

Finally we note the special "type switch" in go.  We can switch on an interface as a tag, and perform different logic against the interface based on its type.
```go
var myInter interface{} = 1
switch myInter.(type) {
case int:
    fmt.Println("myInter is an int") // This will execute
case float64:
    fmt.Println("myInter is a float64")
case string:
    fmt.Println("myInter is a string")
default:
    fmt.Println("myInter is another type")
}
```