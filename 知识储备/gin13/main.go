package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user struct {
	username string
	password string
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(ctx *gin.Context) {
		u := ctx.Query("username")
		p := ctx.Query("password")

		user := user{
			username: u,
			password: p,
		}
		// var u user
		// err := ctx.ShouldBind(&u)
		// if err != nil {
		// 	ctx.JSON(500, gin.H{
		// 		"err": err.Error(),
		// 	})
		// 	return
		// }
		fmt.Printf("%#v\n", user)
		ctx.JSON(200, gin.H{
			"message": "ok",
		})

	})
	r.GET("index",func(ctx *gin.Context) {
		ctx.HTML(200,"index.html",nil)
	})
	r.POST("index", func(ctx *gin.Context) {
		var u user
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
	})
	r.Run(":9001")

}
