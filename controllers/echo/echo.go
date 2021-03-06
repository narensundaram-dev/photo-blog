package echo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
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
