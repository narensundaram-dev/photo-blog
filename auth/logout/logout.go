package logout

import (
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	content := strings.NewReader("You called Logout Page. Method: POST")
	io.Copy(res, content)
}
