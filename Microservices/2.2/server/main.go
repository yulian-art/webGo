package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct {}

func (s *ServiceA) Add (args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}
func main(){
	service := new(ServiceA)
	rpc.Register(service)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9093")
	if e != nil {
		log.Fatal("listen error :", e)
	}
	http.Serve(l, nil)

}