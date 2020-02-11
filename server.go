package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"

	"photo-blog/controllers/index"

	"photo-blog/controllers/auth/login"
	"photo-blog/controllers/auth/logout"
	"photo-blog/controllers/auth/signup"
)

func main() {
	router := getRouter()
	http.ListenAndServe(":8080", router)
}

func getRouter() *httprouter.Router {
	router := httprouter.New()

	// Index Page
	router.GET("/", index.Index)

	// Auth Routers
	router.GET("/signup", signup.Get)
	router.POST("/signup", signup.Post)
	router.GET("/login", login.Get)
	router.POST("/login", login.Post)
	router.POST("/logout", logout.Post)

	return router
}
