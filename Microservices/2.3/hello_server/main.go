package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

// 编写结构体注册服务
type server struct {
	pb.UnimplementedGreeterServer // 当没有完全实现时也可以正常运行
}

const serviceName = "hell0"

// 定义方法
// 方法是对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}


func main(){
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatalf("failed to listen, err: %v", err)
	}

	

	// 注册服务
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// 连接至consul
	cc, err := api.NewClient(api.DefaultConfig()) //8500端口
	if err != nil {
		fmt.Printf("api.NewClient failed:%v\n", err)
		return
	}

	// 将grpc服务注册到consul
	srv := &api.AgentServiceRegistration{
		ID: fmt.Sprintf("%s-%s-%d", serviceName, "127.0.0.1", 8972), // 服务唯一id
		Name: serviceName,
		Tags: []string{"julien"},
		Address: "127.0.0.1",
		Port: 8972,
	}

	if err := cc.Agent().ServiceRegister(srv); err != nil {
		log.Printf("failed to register service, err: %v", err)
		return
	}

	if err := s.Serve(l); err != nil {
		log.Printf("failed to serve, err: %v", err)
	}
}



