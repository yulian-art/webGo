package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("hello failed")
		return
	}
	t.Execute(w, "111")

}

func world(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./world.tmpl")
	if err != nil {
		fmt.Println("world failed")
		return
	}
	t.Execute(w, "11122")

}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("failed")
		return
	}
}
