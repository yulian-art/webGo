package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// data := map[string]any{
		// 	"name" : " julien ",
		// 	"age" : "19",
		// 	"hobby" : "eating",
		// }

		data := gin.H{
			"name":  " julien ",
			"age":   "19",
			"hobby": "eating",
		}
		c.JSON(http.StatusOK, data)
	})

	type msg struct {
		Name  string
		Age   int
		Hobby string
	}

	r.GET("/struct", func(c *gin.Context) {
		data := msg{
			Name:  "julien",
			Age:   18,
			Hobby: "sing",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9000")
}
