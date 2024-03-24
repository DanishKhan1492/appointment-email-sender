package main

import (
	"appointment-notification-sender/main/src/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("view/*")

	router.GET("/", controllers.IndexHandler)
	router.POST("/", controllers.UploadFile)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
		IdleTimeout:  300 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
