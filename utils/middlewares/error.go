package middlewares

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleError is a middleware to respond an error to the client gracefully without any server crash
func HandleError(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

		defer func() {
			if err := recover(); err != nil {
				res.WriteHeader(http.StatusInternalServerError)

				msg := fmt.Sprintf("%v", err)
				res.Write([]byte(msg))
			}
		}()

		handler(res, req, params)
	}
}
