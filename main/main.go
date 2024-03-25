package main

import (
	"appointment-notification-sender/main/src/controllers"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
	"time"
)

//go:embed view/*
var views embed.FS

//go:embed static/*
var staticFiles embed.FS

func main() {
	router := gin.Default()

	// Serve static files
	staticFs, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	router.StaticFS("/static", http.FS(staticFs))

	// Load HTML templates
	tmpl, err := template.ParseFS(views, "view/*.html")
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(tmpl)

	router.GET("/", controllers.IndexHandler)
	uploadHandler := controllers.UploadHandler{TemplateFs: views}
	router.POST("/", func(c *gin.Context) {
		uploadHandler.UploadFile(c)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
		IdleTimeout:  300 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
