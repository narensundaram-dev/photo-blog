package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models"
	"photo-blog/models/db"

	"photo-blog/controllers/index"

	"photo-blog/controllers/auth/login"
	"photo-blog/controllers/auth/logout"
	"photo-blog/controllers/auth/signup"
)

// User Model
type User models.User

// Photo Model
type Photo models.Photo

func main() {
	db := db.OpenDBConnection()
	defer db.Close()

	router := getRouter()
	fmt.Printf("Server opens on port: 8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", router))
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
