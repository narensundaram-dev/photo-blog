package index

import (
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Index Page
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	content := strings.NewReader("This is an index page!")
	io.Copy(res, content)
}
