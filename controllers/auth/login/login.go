package login

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	encrypt "photo-blog/utils"
	jwt "photo-blog/utils"
	response "photo-blog/utils"
	session "photo-blog/utils"
	templates "photo-blog/utils"
)

// User Model
type User models.User

// HTTPResponse responds to the request to render it to form
type HTTPResponse response.Response

// Session to get and set auth token
var Session session.Session = session.Session{}

var tpl *template.Template

func init() {
	tpl = templates.LoadTemplatesGlob()
}

// Get Handler
func Get(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "login.html", req.Form)
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
			tpl.ExecuteTemplate(res, "login.html", response)
		}
	}()

	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")

	user := User{}
	db.Get().First(&user, "username = ?", username)
	if user == (User{}) {
		panic("User not found in the database.")
	}

	isValid := encrypt.VerifyPassword(user.Password, password)
	if !isValid {
		panic("Invalid username/password.")
	}
	j := jwt.Jwt{UID: user.ID, Name: user.Name, Username: user.Username}
	token := j.GenerateToken()
	db.Get().Model(&user).Update("token", token)

	Session.Set(res, token)
	http.Redirect(res, req, "/home", http.StatusSeeOther)
}
