package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is an api route",
		})
	})
	r.GET("/", func(c *gin.Context) {
		data := make(map[string]interface{})
		data["title"] = "Home Page"
		data["body"] = "This is the home page"
		homePage(c,  data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}