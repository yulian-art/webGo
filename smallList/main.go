package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
func initMysql()error{
	dsn := "root:root1234@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open("mysql",dsn)
}

func main() {

	// create database bubble

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLFiles("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(ctx *gin.Context) {

		})
		v1Group.GET("/todo",func(ctx *gin.Context) {

		})
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {

		})
	}
	r.Run()
}