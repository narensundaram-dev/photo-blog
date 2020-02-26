
*[Home](../README.md)*

## Server

**Comprehensive**

``` Go
package main

import (
	"fmt"
	"net/http"
)

// Server to respond to the client request
type Server struct{}

func (server Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Response data here...")
}

func main() {
	server := Server{}
	http.ListenAndServe(":8080", server)
}

```

**Detailed**

- HTTP uses TCP connection
- Server opens up for client request on any port (Typically 80 for HTTP).
- Server has the router component that routes the request to particular part of the code (based on the request).

``` Go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server opens up on port: 8080")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn) // log the request in console
	respond(conn) // respond back to the client
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break // request content is done
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello, World!</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
```

## MVC Architecture:

**Model:** 
- How the data has to be stored (Storing the username and password)

**View:** 
- What the user should see (User inputs the username and password)

**Controller:** 
- Where the business logic has been done (Validating the username and password)