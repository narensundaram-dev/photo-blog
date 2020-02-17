package utils

import (
	"net/http"
	"time"
)

// Session using Cookie
type Session struct{}

// Set Auth Token
func (s *Session) Set(res http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().AddDate(0, 0, 1), // expires in a day
	}
	http.SetCookie(res, cookie)
}

// Get Auth Token
func (s *Session) Get(req *http.Request) string {
	if cookie, err := req.Cookie("token"); err != nil {
		return ""
	} else {
		return cookie.Value
	}
}

// Delete Session Token
func (s Session) Delete(res http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)
}
