package index

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models/db"
	jwt "photo-blog/utils"
	response "photo-blog/utils"
	templates "photo-blog/utils"
)

// HTTPResponse struct
type HTTPResponse response.Response

// Index Page
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl := templates.LoadTemplatesGlob()

	token := Session.Get(req)
	j := &jwt.Jwt{}
	j = j.GetUserInfo(token)

	photos := make([]Photo, 10)
	db.Get().Where("user_id = ?", j.UID).Find(&photos)
	body := make(map[string]interface{})
	body["Photos"] = photos
	body["Name"] = j.Name

	response := HTTPResponse{
		Data: response.Data{
			Body: body,
		},
		Error: response.Error{Message: ""},
	}
	err := tpl.ExecuteTemplate(res, "index.html", response)
	if err != nil {
		panic(err)
	}
}
