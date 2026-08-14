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
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// us := c.PostForm("username")
		// pw := c.PostForm("password")
		us := c.DefaultPostForm("username", "未输入username")

		pw := c.DefaultPostForm("password", "password！！！")
		c.HTML(200, "index.html", gin.H{
			"name": us,
			"pass": pw,
		})
	})
	r.Run(":9000")
}
