package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() error {
	dsn := "root:LYX88888888@tcp(127.0.0.1:3306)/db1"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

func main() {

	// create database bubble
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLFiles("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(ctx *gin.Context) {
			var todo Todo
			ctx.BindJSON(&todo)
			err := DB.Create(&todo).Error
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}
		})
		v1Group.GET("/todo", func(ctx *gin.Context) {
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "err id"})
				return
			}
			var todo Todo
			err = DB.Where("id = ?", id).First(&todo).Error
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			ctx.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": "put err"})

			} else {
				ctx.JSON(http.StatusOK, "ok")
			}
		})
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, _ := ctx.Params.Get("id")
			if err = DB.Where("id = ?", id).Delete(Todo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
			}

		})
	}
	r.Run(":9000")
}
