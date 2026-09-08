package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	X, Y int
}

func main(){
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9093")
	if err !=nil {
		log.Fatal("dailing:", err)
	}

	a:=&Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", a, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error :", err)
	}
	fmt.Printf("ServiceA.Add : %d\n", reply)

	var reply2 int
	divCall := client.Go("ServiceA.Add", a, &reply2, nil)
	replyCall := <- divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}