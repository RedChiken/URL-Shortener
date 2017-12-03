package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/handlers", "./handlers")
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(200, "page.tmpl", gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
