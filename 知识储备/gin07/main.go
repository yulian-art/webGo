package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("hello parse failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, "111"); err != nil {
		fmt.Println("hello execute failed:", err)
	}

}

func world(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./world.tmpl")
	if err != nil {
		fmt.Println("world parse failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, "11122"); err != nil {
		fmt.Println("world execute failed:", err)
	}

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
