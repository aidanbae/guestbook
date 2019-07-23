package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.New()

	r.Static("/css", "./fe/dist/css")
	r.Static("/js", "./fe/dist/js")
	r.Static("/img", "./fe/dist/img")
	r.LoadHTMLFiles("fe/dist/index.html")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "11",
		})
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	authorized.GET("/admin", func(c *gin.Context){
		c.HTML(200, "index.html", nil)
	})

	initRouter(r)

	r.Run(":9000")

}