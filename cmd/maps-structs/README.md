# Maps and Structs

This application explores the various functionality surrounding maps and slices in go.

- [Maps and Structs](#maps-and-structs)
  - [Maps](#maps)
  - [Structs](#structs)

## Maps

Here we explore what maps are in go, how to create them, and how to handle and manipulate their data.

Maps are collection types which aggregate statically typed key-value pairs.  For example, I can declare a map which maps strings to strings, ints to strings, etcetera.
```go
// Example map initialization
myMap := map[string]string {
    "hello" : "world"
}
```

Maps can be created as literals, or they can be created using the built in `make()` function.
```go
// Example map declaration using make
makeMap := make(map[int]string)
```

Map key value pairs are accessed via square bracket syntax like what is seen below.
```go
// Example of accessing map key-value pairs
myMap := map[string]string {
    "hello" : "world"
}
fmt.Println(myMap["hello"]) // world
```

Values corresponding to missing keys in the map are set to the zero value for their type.  We can use "comma-ok" syntax when reading map properties to determine whether a key actually exists within a map.  That is, when accessing a map property, an optional second boolean value is returned identifying whether the key was found in the map.
```go
// Example of comma-ok syntax when accessing map key-value pairs
myMap := map[string]string {
    "hello" : "world"
}

// The "hello" key exists
_, ok := myMap["hello"]
fmt.Println(ok) // true

// The "goodbye" key does not exist
goodbye, ok := myMap["goodbye"]
fmt.Println(ok) // false

// The returned goodbye value is zeroed
fmt.Println(goodbye) // prints empty string "" plus newline
```

Maps are reference types in go, meaning when reassigning a map to a new variable, the new variable is simply a pointer to the original map.
```go
// Example of maps as reference types
myMap := map[string]string {
    "hello" : "world"
}

// Reassigning myMap to new variable
yourMap := myMap

// Mutating yourMap
yourMap["hello"] = "earth"

// Both myMap and yourMap are mutated
fmt.Println(myMap["hello"])   // earth
fmt.Println(yourMap["hello"]) // earth
```

## Structs

Here we explore what structs are in go, how to create them, and the naming conventions surrounding structs in go.  We also explore important concepts around structs such as embedding, and the usage of tags.

Structs are not exactly _collection_ types per se, but they do group disparate data together to describe a single concept.  Structs are keyed by named fields which follow the same syntax as variables in go, including that involving public/private behavior.  Generally, structs are assigned to a type.
```go
// Example of struct declaration and type assignment
type Person struct {
    Name string // public
    age int     // private
}
```

However, structs do not need to be assigned to a type, in which case we have what is called an anonymous struct.  These are rarely used; when they are they are generally short-lived.
```go
// Example of anonymous struct initialization
s := struct{ Name string }{ Name: "Joe" }
fmt.Println(s.Name) // Joe
```

Structs are value types in go, meaning when reassigning a struct to a new variable, the new variable is a copy of the original struct instance.
```go
// Example of structs as value types
myStruct := struct{ Name string }{ Name: "Joe" }
yourStruct := myStruct // Copy happens here
yourStruct.Name = "Bob"
fmt.Println(myStruct.Name)   // Joe
fmt.Println(yourStruct.Name) // Bob
```

Structs support composition using embedding, which give us a limited mechanism for inheritance of properties and some functionality from another struct.  However, this does not result in full polymorphism, as structs which embed another struct are not instances of that other struct.
```go
// Example of struct embedding
type Animal struct {
	Name     string
	SpeedMPH float32
}
type Bird struct {
	Animal     // Embedding - a bird has animal properties
	WingspanCM float32
}
```

In the above example, we note that the `Bird` struct actually has a field named `Animal` which is an instance of the Animal struct containing the values given for its properties when initializing the Bird struct.  However, these properties are aliased in the Bird struct so that they can be referred to as though they were its own properties.
```go
// Example of instantiating a struct using embedding
myBird := Bird{
    // Animal is an explicitly named property
    Animal: Animal{
        // Name and SpeedMPH are sub-properties of the
        // Animal sub-property of the Bird
        Name: "Eurasian Tree Sparrow",
        SpeedMPH: 30.3,
    },
    WingspanCM: 20.2,
}

// But we can access them as though they were direct
// sub-properties of the Bird
fmt.Println(myBird.Name) // Eurasian Tree Sparrow
```

Struct fields can also be tagged with metadata, which can be parsed and leveraged via the `reflect` package.  Generally, tags are used in validation frameworks to ensure that fields adhere to certain arbitrary constraints such as those suggested below, like `required:"true"` or `required:"false"`.
```go
type Dog struct {
	IsGood bool   `required:"false" default:"true" mustSetEqualTo:"true"`
	Breed  string `required:"true" max:"100"`
}
```