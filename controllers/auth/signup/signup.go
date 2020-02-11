package signup

import (
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Get Handler
func Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	content := strings.NewReader("You called Signup page. Method: GET")
	io.Copy(res, content)
}

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	content := strings.NewReader("You called Signup Page. Method: POST")
	io.Copy(res, content)
}
