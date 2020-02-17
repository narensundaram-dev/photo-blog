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
	fmt.Printf("Server opens up on port: 8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter() *httprouter.Router {
	router := httprouter.New()

	// Index Page
	router.GET("/", mw.HandleError(
		mw.IsAuthenticated(index.Index)))

	// Auth Routers
	router.GET("/signup", mw.HandleError(signup.Get))
	router.POST("/signup", mw.HandleError(signup.Post))
	router.GET("/login", mw.HandleError(login.Get))
	router.POST("/login", mw.HandleError(login.Post))
	router.POST("/logout", mw.HandleError(logout.Post))

	// Sample JSON Reponse
	router.Handler("POST", "/api/echo/body", echo.Echo{})
	router.Handler("GET", "/api/echo/body", echo.Echo{})

	// Serve File
	router.GET("/download/file", mw.HandleError(echo.DownloadFile))
	router.POST("/upload/file", mw.HandleError(echo.UploadFile))

	return router
}
