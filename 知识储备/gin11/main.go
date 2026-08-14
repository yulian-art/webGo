package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/qurey", func(c *gin.Context) {
		//获取浏览器发送的querryString
		name := c.Query("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	r.Run(":9000")
}
