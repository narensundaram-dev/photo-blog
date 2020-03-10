*[Home](../README.md)*

## Quick Setup:

- Basic Syntax (Hello, World!)
- Package
- Data Types
- Array & Slice
- Struct & Map
- Variables
- Constants
- Operators
- Decision Making
- Loops
- Functions
- Scope Rules
- Range
- Recursion
- Type Casting
- Interfaces
- Error Handling

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

1. [Numbers](https://tour.golang.org/basics/11)
2. [Strings](https://tour.golang.org/basics/11)
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
type PersonMap map[string]string

func main() {

	personStruct := PersonStruct{
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

## Variables

- Declaring a variable

``` Go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

- Defining a variable

``` Go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

- Python like syntax

``` Go
package main

import "fmt"

func main() {
	k := 3
	fmt.Println(k)
}

```

## Constants

- once defined, can't be redefined (or) updated.

``` Go
package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println(Pi)
}

```

## Operators

- Reference: https://www.tutorialspoint.com/go/go_operators.htm

## Decision Making

- Using if else

``` Go
package main

import (
	"fmt"
	"runtime"
)

var os string = runtime.GOOS

func main() {
	if os == "linux" {
		fmt.Println("Linux OS")
	} else if os == "windows" {
		fmt.Println("Windows OS")
	} else {
		fmt.Println(os)
	}
}

```

- Using switch

``` Go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "nacl":
		fmt.Println("NACL.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

```

## Functions

- A simple number swapping using golang function

``` Go
package main

import (
	"fmt"
)

var x, y int = 10, 20

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Before swap: ", x, y)
	x, y := swap(x, y)
	fmt.Println("After swap: ", x, y)
}

```

## Scope rules

- global variable can be used anywhere within a package
- local variable can be used only inside the function (not outside the function)
- global variable can be override inside any function and can be used as local variable

``` Go
package main

import "fmt"

/* global variable declaration */
var g int = 10

func main() {
	/* local variable declaration */
	var l int = 20
	// var g int = 99 // override global variable

	fmt.Println("value of g and l =", g, l)
}

```

## Range

- Keyword to iterate the iterables

- Iterating Array / Slice

``` Go
package main

import "fmt"

func main() {
	numbers := []string{"One", "Two", "Three"}

	fmt.Println("Iterate by index only:")
	for idx := range numbers {
		fmt.Println("\tSlice index", idx, "is", numbers[idx])
	}

	fmt.Println("\nIterate by index and element")
	for idx, ele := range numbers {
		fmt.Println("\tSlice index", idx, "is", ele)
	}
}

```

- Iterating Map

``` Go
package main

import "fmt"

func main() {

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	fmt.Println("Iterate by key only:")
	for country := range countryCapitalMap {
		fmt.Println("\tCapital of", country, "is", countryCapitalMap[country])
	}

	fmt.Println("\nIterate by key and value:")
	for country, capital := range countryCapitalMap {
		fmt.Println("\tCapital of", country, "is", capital)
	}
}

```

- Iterating Struct (Can't use range. Instead use an alternative.)

``` Go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	person := struct {
		Name   string
		Age    int
		Gender string
	}{"John", 26, "Male"}

	v := reflect.ValueOf(person)
	t := v.Type()

	// fmt.Println("Value: ", v)
	// fmt.Println("Type: ", t)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:\t %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
}

```

## Recursion

- A function calls itself again and again
- To restrict the infinite call, we need a base case

``` Go
package main

import "fmt"

func factorial(i int) int {
	if i <= 1 { // base case that stops recursive call
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	i := 5
	v:= factorial(i)
	fmt.Printf("Factorial of %d is %d\n", i, v)
}

```

## Type Casting

``` Go
```

## Type Assertion

``` Go
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

## Error Handling

### Keywords

- defer
- panic
- recover


## defer

- the execution is postponed until the surrounding code is executed.

``` Go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("I'll execute after below code runs!")
	}()

	x := 1 / 2
	fmt.Println("Divsion: ", x)
}

```

## panic

- to throw an error

``` Go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("I'll execute after all code runs!")
	}()

	x := 4
	y := 0
	if y == 0 {
		panic("y can't be zero!")
	}

	z := x / y
	fmt.Println("Result: ", z)
}

```

## recover

- catches the error and do something with it.

``` Go
package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Exiting gracefuly.")
			os.Exit(1)
		}
	}()

	x := 4
	y := 0
	if y == 0 {
		panic("y can't be zero!")
	}

	z := x / y
	fmt.Println("Result: ", z)
}

```
