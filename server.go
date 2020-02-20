package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"photo-blog/models/db"
	mw "photo-blog/utils/middlewares"

	"photo-blog/controllers/index"

	"photo-blog/controllers/auth/login"
	"photo-blog/controllers/auth/logout"
	"photo-blog/controllers/auth/signup"
	"photo-blog/controllers/echo"
	"photo-blog/controllers/users"
)

func main() {
	// Open DB Connection
	db := db.OpenDBConnection()
	defer db.Close()

	router := getRouter()
	fmt.Printf("Server opens up on port: 8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter() *httprouter.Router {
	router := httprouter.New()

	// Index Page
	router.GET("/", mw.HandleError(
		mw.IsAuthenticated(index.Index)))
	router.GET("/home", mw.HandleError(
		mw.IsAuthenticated(index.Home)))

	// Auth Routers
	router.GET("/signup", mw.HandleError(signup.Get))
	router.POST("/signup", mw.HandleError(signup.Post))
	router.GET("/login", mw.HandleError(login.Get))
	router.POST("/login", mw.HandleError(login.Post))
	router.POST("/logout", mw.HandleError(logout.Post))

	// Upload/Download File
	router.POST("/upload/file", mw.HandleError(
		mw.IsAuthenticated(index.UploadFile)))
	router.GET("/download/file/:id", mw.HandleError(
		mw.IsAuthenticated(index.DownloadFile)))
	router.DELETE("/delete/file/:id", mw.HandleError(
		mw.IsAuthenticated(index.DeleteFile)))
	router.GET("/delete/file/:id", mw.HandleError(
		mw.IsAuthenticated(index.DeleteFile)))

	// Users and their Photos
	router.GET("/users", mw.HandleError(mw.IsAuthenticated(users.Get)))

	// Sample JSON Reponse
	router.Handler("POST", "/api/echo/body", echo.Echo{})
	router.Handler("GET", "/api/echo/body", echo.Echo{})

	// Serving Media Files
	wd, _ := os.Getwd()
	router.ServeFiles("/static/*filepath", http.Dir(wd))

	return router
}
