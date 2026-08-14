package main

import (
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "julien",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a herf='https://julien.net.cn'>julien s blog</a>",
		})
	})
	r.Run(":9000")
}
