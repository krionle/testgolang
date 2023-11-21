package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	rot := gin.Default()
	rot.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "1")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id:%s  page:%s  name:%s  message:%s", id, page, name, message)
	})
	rot.Run(":8080")
}
