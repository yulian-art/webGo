package main

import (

	"github.com/gin-gonic/gin"
)


func sayHello(c *gin.Context){
	c.JSON(200, gin.H{
		"message" : "Hello",
	})
}

func main(){
	//创建默认路由引擎
	r := gin.Default()
	r.GET("/hello",sayHello)
	r.Run()
}