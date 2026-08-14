package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:LYX88888888@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{
	// 	1, "julien", "woman", "eating",
	// }
	// db.Create(&u1)
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	db.Model(&u).Update("hobby", "sleeping")
	fmt.Printf("u:%#v\n", u)
	db.Delete(&u)
}
