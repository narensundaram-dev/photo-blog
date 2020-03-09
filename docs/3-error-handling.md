*[Home](../README.md)*

### Terminologies

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
