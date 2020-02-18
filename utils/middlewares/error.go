package middlewares

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
)

// HandleError is a middleware to respond an error to the client gracefully without any server crash
func HandleError(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

		defer func() {
			if err := recover(); err != nil {
				res.WriteHeader(http.StatusInternalServerError)

				_, fn, line, _ := runtime.Caller(1)
				msg := fmt.Sprintf("Internal Server Error: \n\nLineno: %s:%d\n\nMessage: %v", fn, line, err)
				res.Write([]byte(msg))
			}
		}()

		handler(res, req, params)
	}
}
