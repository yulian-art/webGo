package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"goole.golang.org/grpc"
)

// 编写结构体注册服务
type server struct {
	pb.UnimplementedGreeterServer // 当没有完全实现时也可以正常运行
}

// 定义方法
// 方法是对外提供的服务
func (s *serer) SayHello(ctx context.Context ,in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}


func main(){
	// 启动服务
	l, er := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatal("failed to listen, err :%v\n", err )
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err =  s.Serve(l)
	if err := nil {
		fmt.Printf("failed to serve, err ")
	}
}