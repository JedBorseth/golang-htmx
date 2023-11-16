package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is an api route",
		})
	})
	r.GET("/", func(c *gin.Context) {
		homePage(c)
	})
	r.GET("/style.css", func(c *gin.Context) {
		c.File("styles/output.css")
	})
	// Get all files in the assets folder
	assetsDir := "./assets"
	files, err := os.ReadDir(assetsDir)
	if err != nil {
		log.Fatal(err)
	}
	// Make a GET endpoint for each file
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(assetsDir, file.Name())
			r.GET("/"+file.Name(), func(c *gin.Context) {
				http.ServeFile(c.Writer, c.Request, filePath)
			})
		}
	}
	
	r.Run() // listen and serve on 0.0.0.0:8080
}