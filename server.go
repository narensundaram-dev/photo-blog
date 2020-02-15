package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models/db"
	mw "photo-blog/utils/middlewares"

	"photo-blog/controllers/index"

	"photo-blog/controllers/auth/login"
	"photo-blog/controllers/auth/logout"
	"photo-blog/controllers/auth/signup"
	"photo-blog/controllers/echo"
)

func main() {
	// Open DB Connection
	db := db.OpenDBConnection()
	defer db.Close()

	router := getRouter()
	fmt.Printf("Server opens on port: 8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter() *httprouter.Router {
	router := httprouter.New()

	// Index Page
	router.GET("/", mw.IsAuthenticated(index.Index))

	// Auth Routers
	router.GET("/signup", signup.Get)
	router.POST("/signup", signup.Post)
	router.GET("/login", login.Get)
	router.POST("/login", login.Post)
	router.POST("/logout", logout.Post)

	// Sample JSON Reponse
	router.Handler("POST", "/api/echo/body", echo.Echo{})
	router.Handler("GET", "/api/echo/body", echo.Echo{})

	// Serve File
	router.GET("/serve/file", echo.ServeFile)

	return router
}
