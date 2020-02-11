package index

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	templates "photo-blog/utils"
)

// Index Page
func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl := templates.LoadTemplatesGlob()
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		panic(err)
	}
}
