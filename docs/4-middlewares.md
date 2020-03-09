*[Home](../README.md)*

- A middleware is a piece of code that executes before the request is processed by the handler.
- Eg: To validate the client request is authenticated or not. (i.e) if the request header holds the valid token.

## Function as a 1st class object:

1. Function can be stored in a variable then invoked later.
2. The function variable can also be passed as a parameter to other function (From the function variable can be invoked).


### Justification 1:

``` Go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	mtd := add
	result := mtd(a, b)

	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

```

### Justification 2:

``` Go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	mtd := add
	funcProcessor(mtd, a, b)
}

func add(a, b int) int {
	return a + b
}

func funcProcessor(method func(int, int) int, a, b int) {
	result := method(a, b)
	fmt.Println(result)
}

```

### Apply

``` Go
router.POST("/signup", mw.IsAuthenticated(index.UploadFile))

func IsAuthenticated(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
        // Do something here ...

        // TODO:
        // - Process the request. 
        // - Validate if the header has the valid token.
        // - If not valid token, redirect to login page.
        // - Else just continue with handler function. (i.e), invoke!
    }
}

```