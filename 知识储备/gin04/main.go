package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func sayHello(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err!= nil{
		fmt.Println("template failed: %v",err)
		return
	}
	name := "111"
	err = tmpl.Execute(w, name)
	if err != nil{
		fmt.Println("execulate failed: %v", err)
		return
	}

}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP listen and server failed %v", err)
		return
	}
}