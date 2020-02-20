package users

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	jwt "photo-blog/utils"
	response "photo-blog/utils"
	session "photo-blog/utils"
	templates "photo-blog/utils"
)

// User model
type User models.User

// Session to get and set auth
var Session session.Session = session.Session{}

// HTTPResponse struct
type HTTPResponse response.Response

// Get to see the registered users
func Get(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl := templates.LoadTemplatesGlob()

	token := Session.Get(req)
	j := &jwt.Jwt{}
	j = j.GetUserInfo(token)

	users := make([]User, 0, 10)
	db.Get().Where("id != ?", j.UID).Order("created_at desc").Find(&users)
	body := make(map[string]interface{})
	body["Users"] = users

	response := HTTPResponse{
		Data: response.Data{
			Body: body,
		},
		Error: response.Error{Message: ""},
	}
	err := tpl.ExecuteTemplate(res, "users.html", response)
	if err != nil {
		panic(err)
	}
}
