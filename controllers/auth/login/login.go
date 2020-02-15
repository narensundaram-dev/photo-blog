package login

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	templates "photo-blog/utils"
)

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// Get Handler
func Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(500)
			tpl.ExecuteTemplate(res, "login.html", nil)
		}
	}()

	err := tpl.ExecuteTemplate(res, "login.html", req.Form)
	panic(err)
}

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(500)
			tpl.ExecuteTemplate(res, "login.html", nil)
		}
	}()

	req.ParseForm()
	// username := req.Form.Get("username")
	// password := req.Form.Get("password")
	// fmt.Println(username, password)

	http.SetCookie(res, &http.Cookie{Name: "token", Value: "123"})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
