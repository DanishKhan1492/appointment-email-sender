package controllers

import (
	"appointment-notification-sender/main/src/models"
	"appointment-notification-sender/main/src/utility"
	"embed"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

var indexPage = "index.html"
var customerPage = "customers.html"

type UploadHandler struct {
	TemplateFs embed.FS
}

func (uh UploadHandler) UploadFile(c *gin.Context) {
	log.Print("Request Received")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.HTML(http.StatusBadRequest, indexPage, gin.H{"error": "Please upload a file"})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(file)

	log.Println("Reading the File")
	// Read the uploaded file into memory
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.HTML(http.StatusInternalServerError, indexPage, gin.H{"error": "Error reading file"})
		return
	}

	// Check file extension to determine if it's CSV or XLSX
	ext := filepath.Ext(header.Filename)
	log.Println("File is of type : " + ext)
	var customers []models.Customer
	if ext == ".csv" {
		// Parse CSV file directly from memory
		customers, err = utility.ParseCSVFromMemory(fileBytes)
	} else if ext == ".xlsx" {
		// Parse XLSX file directly from memory
		customers, err = utility.ParseXLSXFromMemory(fileBytes)
	} else {
		c.HTML(http.StatusBadRequest, indexPage, gin.H{"error": "Unsupported file format"})
		return
	}

	if err != nil {
		c.HTML(http.StatusInternalServerError, indexPage, gin.H{"error": "Error parsing file"})
		return
	}

	// Send messages to customers
	utility.SendMessages(&customers, uh.TemplateFs)

	// Render results to HTML
	c.HTML(http.StatusOK, customerPage, gin.H{"customers": customers})
}
