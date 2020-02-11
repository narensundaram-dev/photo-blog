package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"photo-blog/index"

	"photo-blog/auth/login"
	"photo-blog/auth/logout"
	"photo-blog/auth/signup"
)

func main() {
	router := httprouter.New()

	// Index Page
	router.GET("/", index.Index)

	// Auth Routers
	router.GET("/signup", signup.Get)
	router.POST("/signup", signup.Post)
	router.POST("/login", login.Post)
	router.POST("/logout", logout.Post)

	// Opening up the port 8080 to listen & respond to the outside request
	http.ListenAndServe(":8080", router)
}
