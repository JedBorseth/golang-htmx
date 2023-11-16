package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context, content any) {
	templ, err := template.ParseFiles("templates/header.html")
		if(err != nil) {
		fmt.Print("Error parsing row template\n ", err)
		return
		}
		templ.Execute(c.Writer, content)
}