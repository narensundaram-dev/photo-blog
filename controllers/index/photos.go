package index

import (
	"fmt"
	"io"
	"net/http"
	"os"
	fp "path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"
	jwt "photo-blog/utils"
	photos "photo-blog/utils"
	session "photo-blog/utils"
)

// Photo Model
type Photo models.Photo

// Session to get and set auth token
var Session session.Session = session.Session{}

// UploadFile to get file from request
func UploadFile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.ParseMultipartForm(1024 * 1024 * 5) // Max size 5242880 bytes. (i.e) 5 MB

	f, h, e := req.FormFile("photo")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	mediaPath := photos.GetMediaPath()
	token := Session.Get(req)
	j := &jwt.Jwt{}
	j = j.GetUserInfo(token)

	userPath := fp.Join(mediaPath, j.Username)
	filePath := fp.Join(mediaPath, j.Username, h.Filename)
	relativePath := photos.GetRelativeMediaPath(j.Username, h.Filename)
	photo := Photo{
		Path:         filePath,
		RelativePath: relativePath,
		UserID:       j.UID,
		Caption:      req.FormValue("description"),
	}
	db.Get().Create(&photo)

	photos.MakeSureThePath(userPath)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, f)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

// DownloadFile to send file as response
func DownloadFile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	file, err := os.Open("/home/naren/Desktop/notes.txt")
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
	} else {
		defer file.Close()

		fnameArr := strings.Split(file.Name(), "/")
		fname := fnameArr[len(fnameArr)-1]
		dispose := fmt.Sprintf(`attachment; filename="%s"`, fname)
		res.Header().Set("Content-Disposition", dispose)
		res.Header().Set("Content-Type", "text/plain")

		fstat, _ := file.Stat()
		size := fmt.Sprintf("%v", fstat.Size())
		res.Header().Set("Content-Length", size)

		res.WriteHeader(http.StatusOK)
		http.ServeFile(res, req, file.Name())
	}

}
