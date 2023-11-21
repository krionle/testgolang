package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.LoadHTMLFiles("./zh.html")
	c.GET("/testhtml", func(r *gin.Context) {
		r.HTML(200, "zh.html", gin.H{
			"title": "main website",
		})
	})
	c.Run(":8080")
}
