package signup

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	templates "photo-blog/utils"
)

// User Model
type User models.User

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// Get Handler
func Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(500)
			msg := fmt.Sprintf("Something's wrong: %s", err)
			io.Copy(res, strings.NewReader(msg))
		}
	}()

	err := tpl.ExecuteTemplate(res, "signup.html", req.Form)
	panic(err)
}

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(500)
			fmt.Printf("Err: %v", fmt.Sprintf("%v", err))
			tpl.ExecuteTemplate(res, "signup.html", nil)
		}
	}()

	// Get form values
	req.ParseForm()
	name := req.Form["name"][0]
	username := req.Form["username"][0]
	password := req.Form["password"][0]

	result := db.Get().Create(&User{
		Name:     name,
		Username: username,
		Password: password,
	})
	if err := result.Error; err != nil {
		// msg := err.Error()
		tpl.ExecuteTemplate(res, "signup.html", nil)
	} else {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}
