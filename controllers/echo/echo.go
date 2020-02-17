package echo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	fp "path/filepath"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"

	photos "photo-blog/utils"
)

// Echo Restful API - Sample Web Service
type Echo struct {
	Message        string
	AnotherMessage string
	PingTime       time.Time
}

func (e Echo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	method := req.Method
	fmt.Printf("Method: %v", method)
	switch method {
	case "GET":
		e.Get(req, res)
	case "POST":
		e.Post(req, res)
	default:
		e.Error(res, http.StatusMethodNotAllowed, "")
	}
}

// Get handler for the request of method GET
func (e *Echo) Get(req *http.Request, res http.ResponseWriter) {
	log.Print("Implement your GET handler here.")
	res.WriteHeader(http.StatusNotImplemented)
}

// Post handler for the request of method POST
func (e *Echo) Post(req *http.Request, res http.ResponseWriter) {
	if err := e.ParseBody(req, res); err != nil {
		e.Error(res, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		if body, err := json.Marshal(e); err != nil {
			e.Error(res, http.StatusBadRequest, fmt.Sprintf("%v", err))
		} else {
			res.WriteHeader(http.StatusOK)
			res.Write(body)
		}
	}
}

// ParseBody from request
func (e *Echo) ParseBody(req *http.Request, res http.ResponseWriter) error {
	e.PingTime = time.Now().Local()
	if err := json.NewDecoder(req.Body).Decode(e); err != nil {
		msg := fmt.Sprintf("%v", err)
		return errors.New(msg)
	}
	return nil
}

func (e *Echo) Error(res http.ResponseWriter, status int, message string) {
	res.WriteHeader(http.StatusBadRequest)

	err := struct{ Message string }{message}
	body, _ := json.Marshal(err)
	res.Write(body)
}

// DownloadFile to send file as response
func DownloadFile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	file, err := os.Open("/home/naren/Desktop/show.txt")
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
	} else {
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
	defer file.Close()

}

// UploadFile to get file from request
func UploadFile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.ParseMultipartForm(1024 * 1024 * 5) // Max size 5242880 bytes. (i.e) 5 MB

	photo, header, err := req.FormFile("photo")
	if err != nil {
		panic(err)
	}
	defer photo.Close()

	mediaPath, err := photos.GetMediaPath()
	if err != nil {
		panic(err)
	}
	filePath := fp.Join(mediaPath, header.Filename)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, photo)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
