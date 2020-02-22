package index

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models/db"
	jwt "photo-blog/utils"
	paginator "photo-blog/utils"
	response "photo-blog/utils"
	templates "photo-blog/utils"
)

// HTTPResponse struct
type HTTPResponse response.Response

// PageLimit limits the records per page
var PageLimit = 6

// Index Page
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(res, req, "/home", http.StatusSeeOther)
}

// Home Page
func Home(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl := templates.LoadTemplatesGlob()

	token := Session.Get(req)
	j := &jwt.Jwt{}
	j = j.GetUserInfo(token)

	page := 1
	pageStr := req.URL.Query().Get("page")
	if pageStr == "" {
		page = 1
	} else {
		pageNo, err := strconv.Atoi(pageStr)
		page = pageNo
		if err != nil {
			panic(err)
		}
	}
	count := 0
	db.Get().Model(&Photo{}).Where("user_id = ?", j.UID).Count(&count)
	pagination := paginator.Paginate(page, PageLimit, count)

	photos := make([]Photo, 0, 10)
	db.Get().Where("user_id = ?",
		j.UID).Offset(pagination.Offset).Order("created_at desc").Limit(PageLimit).Find(&photos)
	body := make(map[string]interface{})
	body["Photos"] = photos
	body["Name"] = j.Name

	body["Previous"] = pagination.Previous
	body["Next"] = pagination.Next
	body["PreviousPageNo"] = page - 1
	body["NextPageNo"] = page + 1

	response := HTTPResponse{
		Data: response.Data{
			Body: body,
		},
		Error: response.Error{Message: ""},
	}
	err := tpl.ExecuteTemplate(res, "home.html", response)
	if err != nil {
		panic(err)
	}
}
