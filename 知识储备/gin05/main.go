package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse failed :%v", err)

	}
	u1 := User{
		Name:   "yyy",
		gender: "woman",
		Age:    19,
	}
	m1 := map[string]any{
		"name": "lyc",
		"age":  1333,
	}
	t.Execute(w, map[string]any{
		"u1": u1,
		"m1": m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("listen and server err :%v", err)
		return
	}

}
