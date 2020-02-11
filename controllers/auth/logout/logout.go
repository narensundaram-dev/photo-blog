package logout

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
