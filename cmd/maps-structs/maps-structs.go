package main

import (
	"fmt"
	"reflect"
)

/*
Creation of structs

Here we declare a type definition, which is
generally a collection of existing types in go.

We note the naming conventions used to define
this struct.  Our struct itself is named with a
variable name starting with a capital letter.
This means our struct is exposed externally.

Meanwhile, most of our struct properties are named
using a leading lower-case character.  This means
that our struct's properties are private.

If we would like to publicly expose our struct's
properties, then we may name them using a capital
leading character.

Note that this Person struct is instantiated and
manipulated below.
*/
type Person struct {
	Name       string   // Public
	birthDay   int      // Private
	birthMonth int      // Private
	interests  []string // Private
}

/*
Struct embedding

Here we define some new Struct types to demonstrate
embedding and composition in go.  It is the case
that go does not support traditional object-oriented
inheritance.  It instead uses embedding and composition.

Here, we want to say that a bird is an animal.  But
go does not support inheritance like most object
oriented languages tend to.  Instead, we use embedding
to say that a bird has all the properties of an animal.

Note that this does not fully solve the problem.  This
gives us the ability to quickly and easily declare
that a bird has all the properties of an animal.  But
it does not give us a way to define a bird as a subtype
of the larger animal type.

In order to do this, we will later explore interfaces
in go.

Embedding in go makes the most sense when we simply want
behavior of a complex base type to be carried forward
into simpler implementations on top of those base types.
It does not make sense when we want a sub-type to be
used interchangably with instances of its parent types.
*/
type Animal struct {
	Name     string
	SpeedMPH float32
}

type Bird struct {
	Animal     // Embedding - a bird has animal properties
	WingspanCM float32
}

/*
Struct tagging

Here we define a new Struct type to demonstrate tagging
in go.  Tagging is generally used to communicate metadata
about your struct to consumers of the struct, say if it
is exposed in a library.

We see that for our dog struct, we specified tags for its
properties.  These tags are given as space-separated key-
value pairs wrapped in backticks.  They are given as
`key1:"value1" key2:"value2"`.

We use the built-in reflect package in go to work with
tags, which we explore below.

Generally, tags are used in validation frameworks to
ensure that fields adhere to certain arbitrary constraints
such as those suggested below, like required:"true" or
required:"false".
*/
type Dog struct {
	IsGood bool   `required:"false" default:"true" mustSetEqualTo:"true"`
	Breed  string `required:"true" max:"100"`
}

func main() {
	/*
		Creation of maps

		Maps are types which aggregate statically-typed
		key-value pairs.  For a type to be used as a key,
		they must be testable for equality.  Slices, maps,
		and functions lack equivalence relations and thus
		cannot be used as keys in maps.

		The built-in make function can be used to declare
		a map.  This is popularly used when the entries of
		the map are not available before compile time and
		thus it must be populated dynamically.
	*/
	fmt.Println("#### Creation of maps ####")

	// An example of a map containing string: int
	// key-value pairs
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Printf("statePopulations: (%T) %v\n", statePopulations, statePopulations)

	// An example of a map containing slice: string
	// key-value pairs - this will fail as slices do
	// not have an equivalence relation defined.
	//sliceMap := map[[]int]string{}
	//fmt.Println(sliceMap)

	// An example of a map containing array: string
	// key-value pairs - this will pass as arrays do
	// have an equivalence relation defined.
	arrayMap := map[[3]int]string{}
	fmt.Printf("arrayMap: (%T) %v\n", arrayMap, arrayMap)

	// An example of a map declared via the make function
	makeMap := make(map[string]int)
	fmt.Printf("makeMap: (%T) %v\n", makeMap, makeMap)
	fmt.Println("")

	/*
		Handling maps & manipulating map data

		To read a value from a map, we may supply a value
		in square brackets equivalent to a key of the map.
		This returns that key's corresponding value.

		To insert a value into a map, we may similarly pass
		a new key in square brackets and assign a value to
		it via the = operator.

		Note that maps are not guaranteed a particular key
		order.  Thus we must be careful when iterating through
		a map, as its keys may be sorted according to some
		unexpected ordering.

		To delete a value from a map, we use the built-in
		delete function.  This function accepts the map as
		its firt argument, and the key to delete as its second
		argument.

		When attempting to read a value from a map using a
		key that does not exist, the resulting value is zeroed
		and thus can cause confusion as to whether it is a
		legitimate key-value pair or whether the key is simply
		missing from the map.

		To interrogate this condition, we use the "comma-ok"
		pattern to accept an optional boolean return value
		which returns false when the key is missing from the
		map, and true when the key exists within the map.

		We may also use the write-only "_" operator to run
		existence checks only without retrieving the value.

		We may use the built-in len() function to determine
		the number of key-value pairs in the map.

		Lastly, we note that maps are reference types in go.
		Thus, when reassigning maps to a new variable, the
		new variable points back at the same map on the heap.
	*/
	fmt.Println("#### Handling maps & manipulating map data ####")

	// Example of reading a value from a map
	ohioPopulation := statePopulations["Ohio"]
	fmt.Printf("ohioPopulation: (%T) %v\n", ohioPopulation, ohioPopulation)

	// Example of inserting a value into a map
	// Example of map ordering changing when value is inserted
	fmt.Printf("statePopulations: (%T) %v\n", statePopulations, statePopulations)
	statePopulations["Georgia"] = 10310371 // Inserted somewhere in middle of map
	georgiaPopulation := statePopulations["Georgia"]
	fmt.Printf("georgiaPopulation: (%T) %v\n", georgiaPopulation, georgiaPopulation)
	fmt.Printf("statePopulations: (%T) %v\n", statePopulations, statePopulations)

	// Example of deleting a value from a map
	delete(statePopulations, "Georgia")
	fmt.Printf("statePopulations: (%T) %v\n", statePopulations, statePopulations)

	// Example of "comma-ok" error check when reading
	// nonexistent key-value pair
	northCarolinaPopulation, ok := statePopulations["North Carolina"]
	fmt.Printf("northCarolinaPopulation: (%T) %v\n", northCarolinaPopulation, northCarolinaPopulation)
	fmt.Printf("Did key exist in map? %v\n", ok)

	// Example of "comma-ok" error check when reading
	// existent key-value pair
	pennsylvaniaPopulation, ok := statePopulations["Pennsylvania"]
	fmt.Printf("pennsylvaniaPopulation: (%T) %v\n", pennsylvaniaPopulation, pennsylvaniaPopulation)
	fmt.Printf("Did key exist in map? %v\n", ok)

	// Example of "comma-ok" error check using the write
	// only operator
	_, ok = statePopulations["Ohio"]
	fmt.Printf("Did key exist in map? %v\n", ok)

	// Example of len() function to determine map length
	fmt.Printf("len(statePopulations): %v\n", len(statePopulations))

	// Example of map reassignment behavior
	newMap := statePopulations
	delete(newMap, "Ohio") // Deletes from both newMap and statePopulations
	fmt.Printf("statePopulations: (%T) %v\n", statePopulations, statePopulations)
	fmt.Printf("newMap: (%T) %v\n", newMap, newMap)
	fmt.Println("")

	/*
		Creation of structs (cont.)

		Here we instantiate a variable of type person, which
		is the struct we previously defined.

		We may also implicitly define struct properties using
		their position within the struct; that is "positional
		syntax".  However this is not maintainable and is
		generally advised against.

		When structs are initialized without certain properties
		having been initialized, those properties are zeroed
		by default.

		Here we also explore the creation of an anonymous
		struct.  We do not need to assign a struct definition
		to a type, instead we may define an initialize it
		directly, resulting in an anonymous struct.

		Anonymous structs are generally very short-lived, and
		are only used when data needs to be grouped in some
		unsupported way via another collection type.
	*/
	fmt.Println("#### Creation of structs ####")

	// Example of instantiating a struct
	myPerson := Person{
		Name:       "Joe",
		birthDay:   21,
		birthMonth: 12,
		interests: []string{
			"Programming",
			"Mathematics",
		},
	}
	fmt.Printf("myPerson: (%T) %v\n", myPerson, myPerson)

	// Example of instantiating a struct using positional
	// syntax (not recommended)
	posPerson := Person{
		"Dave",
		11,
		22,
		[]string{
			"Golf",
			"Football",
		},
	}
	fmt.Printf("posPerson: (%T) %v\n", posPerson, posPerson)

	// Example of struct with zeroed properties
	zeroPerson := Person{}
	fmt.Printf("zeroPerson: (%T) %v\n", zeroPerson, zeroPerson)

	// Example of anonymous struct
	anonStruct := struct{ name string }{name: "Anonymous"}
	fmt.Printf("zeroPerson: (%T) %v\n", anonStruct, anonStruct)
	fmt.Println("")

	/*
		Handling structs and manipulating struct data

		We may access struct data via the dot syntax.  That
		is, we say myVal := myStruct.structVal.  Using this
		dot syntax, we can drill down into nested collections
		within the struct.

		We note that structs are not reference types, and thus
		we must explicitly point back to structs upon reassign-
		ment to mutate the original.
	*/
	fmt.Println("#### Handling structs and manipulating struct data ####")

	// Example of reading a struct property
	fmt.Printf("myPerson.Name: (%T) %v\n", myPerson.Name, myPerson.Name)

	// Example of drilling down into struct data
	fmt.Printf("myPerson.interests: (%T) %v\n", myPerson.interests, myPerson.interests)
	fmt.Printf("myPerson.interests[0]: (%T) %v\n", myPerson.interests[0], myPerson.interests[0])

	// Example of structs as value types
	yourPerson := myPerson // Copy happens here
	yourPerson.Name = "Hank"
	fmt.Printf("myPerson.Name: %v\n", myPerson.Name)
	fmt.Printf("yourPerson.Name: %v\n", yourPerson.Name)

	// Example of using pointers to mutate original struct
	theirPerson := &myPerson // Pointer to original struct
	theirPerson.Name = "Bob"
	fmt.Printf("myPerson.Name: %v\n", myPerson.Name)
	fmt.Printf("theirPerson.Name: %v\n", theirPerson.Name)
	fmt.Println("")

	/*
		Embedding demonstration

		Here we instantiate our bird struct which embeds
		the properties of the animal struct in its
		definition.

		We then explore a bit deeper into the resulting
		internal structure of the struct on which we
		applied the embedding.
	*/
	fmt.Println("#### Embedding demonstration ####")

	// Instantiate a bird which embeds the animal struct
	myBird := Bird{}
	myBird.Name = "Northern Flicker"
	myBird.WingspanCM = 51.0
	myBird.SpeedMPH = 40.5
	fmt.Printf("myBird: (%T) %v\n", myBird, myBird)

	// Reading bird struct property
	fmt.Printf("myBird.WingspanCM: (%T) %v\n", myBird.WingspanCM, myBird.WingspanCM)

	// Reading embedded animal struct property
	fmt.Printf("myBird.Name: (%T) %v\n", myBird.Name, myBird.Name)

	// Instantiate a bird using explicit property names
	expBird := Bird{
		// Notice the Animal embedding is actually an internal
		// property
		Animal: Animal{
			Name:     "Eurasian Tree Sparrow",
			SpeedMPH: 30.3,
		},
		WingspanCM: 20.2,
	}
	fmt.Printf("expBird: (%T) %v\n", expBird, expBird)

	// Demonstrating that the internal animal property exists
	// as a result of embedding, its sub-properties are aliased
	// in the bird struct.
	fmt.Printf("expBird.Animal: (%T) %v\n", expBird.Animal, expBird.Animal)
	fmt.Println("")

	/*
		Struct tagging demonstration

		Here we explore our Dog struct's tags via the reflect
		package.
	*/
	fmt.Println("#### Struct tagging demonstration ####")

	// Using the reflect package to explore our struct's
	// property's tags
	t := reflect.TypeOf(Dog{})
	field, _ := t.FieldByName("Breed")
	fmt.Println(field.Tag)
}
