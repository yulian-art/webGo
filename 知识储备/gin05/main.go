package main

import (
	"fmt"
	"net/http"
	"html/template"
)
type User struct{
	Name string
	gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request){
	t, err:=template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse failed :%v", err)

	}
	u1 := User{
		Name: "yyy",
		gender: "woman",
		Age: 19,
	}
	t.Execute(w, u1)
}

func main() {
	http.HandleFunc("/", sayHello)
	err:=http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("listen and server err :%v", err)
		return
	}


}