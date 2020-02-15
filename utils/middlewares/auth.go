package middlewares

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	session "photo-blog/utils"
	templates "photo-blog/utils"
)

// User Model
type User models.User

// Session to get and set auth token
var Session session.Session = session.Session{}

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// IsAuthenticated is a middleware to check whether the requested user is whether authenticated or not.
func IsAuthenticated(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()

		token := Session.Get(req)
		if token == "" {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		} else {
			handler(res, req, params)
		}
	}
}
