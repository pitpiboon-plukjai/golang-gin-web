package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})

}

func nextPage(c *gin.Context) {

	c.HTML(http.StatusOK, "next.html", gin.H{
		"title": "Second Page",
	})

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexPage)
	r.POST("/next", nextPage)
	r.GET("/next.html", nextPage)

	r.Run()

}
