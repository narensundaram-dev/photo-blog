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
	response "photo-blog/utils"
	templates "photo-blog/utils"
)

// User Model
type User models.User

// HTTPResponse responds to the request to render it to form
type HTTPResponse response.Response

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// Get Handler
func Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("Something's wrong loading the page: %s", err)
			io.Copy(res, strings.NewReader(msg))
		}
	}()

	err := tpl.ExecuteTemplate(res, "signup.html", req.Form)
	panic(err)
}

// Post Handler
func Post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Handle the error gracefully
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%v", err)
			response := HTTPResponse{
				Data:  response.Data{},
				Error: response.Error{Message: msg},
			}
			tpl.ExecuteTemplate(res, "signup.html", response)
		}
	}()

	// Get form values
	req.ParseForm()
	name := req.Form["name"][0]
	username := req.Form["username"][0]
	password := req.Form["password"][0]

	// Push it to DB
	result := db.Get().Create(&User{
		Name:     name,
		Username: username,
		Password: password,
	})
	if err := result.Error; err != nil {
		panic(err)
	} else {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
	}
}
