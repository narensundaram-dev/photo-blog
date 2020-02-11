package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HTTPResponse struct
type HTTPResponse struct {
	Error string
	Data  interface{}
}

// Respond is a decorator that responds to the client
func Respond(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				res.WriteHeader(500)
				msg := fmt.Sprintf("%v", err)
				io.Copy(res, strings.NewReader(msg))
			}
		}()

		handler(res, req)
	})
}
