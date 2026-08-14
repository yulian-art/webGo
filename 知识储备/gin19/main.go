package main

import (
	"database/sql"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string `gorm:"type:varchar(120);unique_index"`
	Role         string `gorm:"size:255"`
	MemberNumber string `gorm:"unique;nut null"`
	Num          int    `gorm:"AUTO_INCREMENT"`
	Address      string `gorm:"index:addr"`
	IgnoreMe     int    `gorm:"-"`
}
type Animals struct{
	AniID int `gorm:"primaty_key"`
	Age int
	Type string
}

func main() {
	db, err := gorm.Open("mysql", "root:LYX88888888@(127.0.0.1:3306)/db1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animals{})

}
