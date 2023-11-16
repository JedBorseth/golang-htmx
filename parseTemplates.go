package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)
 
func header(c *gin.Context, content any) {
	templ, err := template.ParseFiles("templates/header.html")
		if(err != nil) { 
		fmt.Print("Error parsing row template\n ", err)
		return
		} 
		templ.Execute(c.Writer, content)
}

func homePage(c *gin.Context ) {
	data := make(map[string]interface{})
	buf := new(bytes.Buffer)
	data["title"] = "Home Page"

	templ, err := template.ParseFiles("templates/home.html")
		if(err != nil) {
		fmt.Print("Error parsing row template\n ", err)
		return
		}
	templ.Execute(buf, data)
	data["body"] = template.HTML(buf.String())
	header(c, data)
	
}