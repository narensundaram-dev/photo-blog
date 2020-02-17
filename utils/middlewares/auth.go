package middlewares

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	jwt "photo-blog/utils"
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
		token := Session.Get(req)
		if token == "" {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
		} else {
			user := User{}
			db.Get().First(&user, "token = ?", token)
			if user == (User{}) {
				http.Redirect(res, req, "/login", http.StatusSeeOther)
			}
			j := jwt.Jwt{UID: user.ID, Name: user.Name, Username: user.Username}
			isValid := j.ValidateToken(token)
			if !isValid {
				http.Redirect(res, req, "/login", http.StatusSeeOther)
			}
			handler(res, req, params)
		}
	}
}
