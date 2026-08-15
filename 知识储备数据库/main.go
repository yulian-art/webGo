package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMySQL()(err error){
	dsn := "root:LYX88888888@tcp(127.0.0.1:3306)/db1"
	db, err = sql.Open("mysql", dsn)
	if err!= nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("connect failed %v \n", err)
		return
	}

	db.SetConnMaxLifetime(time.Second * 10)
	db.SetConnMaxIdleTime(200)
	db.SetMaxOpenConns(10)

	return
}

type User struct{
	id int
	age int
	name string
}

func queryRowDemo(){
	sqlStr := "select id, name, age from user where id=?"

	var u User

	row:= db.QueryRow(sqlStr, 1)
	
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("Scan failed %v \n",err)
		return
	}
	fmt.Printf("id : %d name %s age : %d", u.id, u.name, u.age)

}

func queryMultiRowDemo(){
	sqlStr := "select id, name, age from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed , err %v \n", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed ,err: %v\n", err)
			return
		}
		fmt.Printf("id : %d, name: %s, age : %d", u.id, u.name, u.age)
	}
}

func main(){
	if err := initMySQL(); err != nil{
		fmt.Printf("connect to db failed , err : %v \n", err)
	}
	defer db.Close()
	queryRowDemo()
	queryMultiRowDemo()

}