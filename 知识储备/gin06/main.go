package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	kua := func(name string) (string, error) {
		return name + "shuai", nil
	}
	t := template.New("hello.tmpl")
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	_, err := t.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("failed")
		return
	}

	t.Execute(w, "111")
}

func tempDemo(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("parse failed")
	}
	t.Execute(w, "111")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/temp", tempDemo)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("failed")
		return
	}
}
