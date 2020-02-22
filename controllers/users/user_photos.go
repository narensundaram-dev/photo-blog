package users

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	paginator "photo-blog/utils"
	response "photo-blog/utils"
	templates "photo-blog/utils"
)

// Photo Model
type Photo models.Photo

// PageLimit limits the records per page
var PageLimit = 6

// UserPhotos to see the particular user uploads
func UserPhotos(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userName := params.ByName("username")

	user := User{}
	db.Get().Where("username = ?", userName).First(&user)

	count := 0
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

	db.Get().Model(&Photo{}).Where("user_id = ?", user.ID).Count(&count)
	pagination := paginator.Paginate(page, PageLimit, count)

	photos := make([]Photo, 0, 10)
	db.Get().Where("user_id = ?",
		user.ID).Offset(pagination.Offset).Order("created_at desc").Limit(PageLimit).Find(&photos)

	body := make(map[string]interface{})
	body["Photos"] = photos
	body["Name"] = user.Name
	body["UserName"] = user.Username

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
	tpl := templates.LoadTemplatesGlob()
	err := tpl.ExecuteTemplate(res, "user_photos.html", response)
	if err != nil {
		panic(err)
	}
}
