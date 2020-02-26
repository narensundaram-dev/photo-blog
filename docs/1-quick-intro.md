*[Home](../README.md)*

## Quick Setup:

- Hello, World!
- Package
- Data Types
- Interface

## Hello, World!

``` Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Package

- A Package is a folder (of one or many *.go files)
- A Root package is a main package
- A sub-package is a sub folder
- No main, No run
- One package, One name
- [Third Party Packages](https://awesome-go.com)
- [Good Read: Medium](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)

## Data Types

1. Numbers (Refer Go Tour)
2. Strings (Refer Go Tour)
3. Array and Slice (Index based)
4. Struct and Map (Key-Value based)

**3. Array & Slice**

|Array|Slice  |
|--|--|
| Value based | Reference based. ([A Good Pic](https://drive.google.com/open?id=1HOmG6cRUL3XFlZxl6G-E9BGXFfu6J0xL)) |
| Fixed size | Dynamically sized |
| Use only when needed | Very commonly used |
| Holds value | Holds reference. (i.e) Slice has 3 properties. Length, Capacity and Pointer to the underlying array |

``` Go
package main

import  "fmt"

func  main() {
	arr  := [3]string{"Naren", "HCL", "Chennai"} // Value based
	slc  :=  make([]string, 0, 3) // Reference based


	fmt.Println("Fixed array: ", arr)
	fmt.Println("Before defining slice: ", slc, len(slc), cap(slc))

	slc  =  append(slc, "Naren")
	slc  =  append(slc, "HCL", "Elcot")
	slc  =  append(slc, "Sholinganallur", "Chennai")
	
	fmt.Println("After defining slice: ", slc, len(slc), cap(slc))

}
```

**4. Struct & Map**

|Struct|Map  |
|--|--|
| Value based | Reference based |
| Fixed-size / Pre-defined | Dynamically sized |
| Very commonly used | Use only when needed |

``` Go
package main

import  "fmt"

// PersonStruct struct
type  PersonStruct  struct {
	Name string
	Place string
}

// PersonMap map
type  PersonMap  map[string]string

func  main() {

	personStruct  := PersonStruct{
		Name: "Naren",
		Place: "Chennai",
	}

	personMap  := PersonMap{}
	personMap["Name"] =  "Naren"
	personMap["Place"] =  "Chennai"
	personMap["Gender"] =  "Male"

	fmt.Println("Before Update: ")
	fmt.Printf("Struct => %+v\n", personStruct)
	fmt.Println("Map => ", personMap)

	personMap.updateName("Narendran G")
	personStruct.updateName("Narendran G")

	fmt.Println("\nAfter Update: ")
	fmt.Printf("Struct => %+v\n", personStruct)
	fmt.Println("Map => ", personMap)
}

func (person PersonMap) updateName(name string) {
	person["Name"] = name
}

func (person *PersonStruct) updateName(name string) {
	person.Name  = name
}
```

## Interface

- For Polymorphism
- To create the generic code

``` Go
package main

import  "fmt"

type  bot  interface {
	getGreeting() string
}

type  englishBot  struct{}

type  tamilBot  struct{}

func  main() {

	eb  := englishBot{}
	tb  := tamilBot{}

	printGreeting(eb)
	printGreeting(tb)
}

func  printGreeting(b bot) {
	// Generic functionality
	greeting  := b.getGreeting()
	fmt.Println(greeting)
}

func (englishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return  "Hello!"
}

func (tamilBot) getGreeting() string {
	// Custom logic for generating a tamil greeting
	return  "Vanakkam!"
}
```