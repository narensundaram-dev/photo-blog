package users

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserPhotos to see the particular user uploads
func UserPhotos(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Println(res, req, params)
}
