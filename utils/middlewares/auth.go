package middlewares

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	templates "photo-blog/utils"
)

// User Model
type User models.User

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// IsAuthenticated is a middleware to check whether the requested user is whether authenticated or not.
func IsAuthenticated(handler httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()

		token, err := req.Cookie("token")
		if err != nil {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		} else if token.Value == "123" {
			handler(res, req, params)
		} else {
			panic("Invalid token")
		}
	}
}
