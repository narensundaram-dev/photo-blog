package logout

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	session "photo-blog/utils"
)

// Session to set and get cookie
var Session session.Session = session.Session{}

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	Session.Delete(res)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
