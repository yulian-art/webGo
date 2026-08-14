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
	r.GET("/user", func(ctx *gin.Context) {
		u := ctx.Query("username")
		p := ctx.Query("password")

		user := user{
			username: u,
			password: p,
		}

		fmt.Printf("%#v\n", user)
		ctx.JSON(200, gin.H{
			"message": "ok",
		})

	})
	r.Run(":9001")

}
