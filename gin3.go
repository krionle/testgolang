package main

import "github.com/gin-gonic/gin"

func main() {
	ht := gin.Default()
	//ht.LoadHTMLFiles("zh.html", "zh-2.html") 读取单个文件
	//读取文件夹
	ht.LoadHTMLGlob("./zh*")

	ht.GET("/t1", func(c *gin.Context) {
		c.HTML(200, "zh.html", gin.H{})
	})
	ht.GET("/t2", func(c *gin.Context) {
		c.HTML(200, "zh-2.html", gin.H{})
	})
	ht.Run(":8080")
}
