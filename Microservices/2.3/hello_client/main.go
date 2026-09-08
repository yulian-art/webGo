package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"hello_client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "julien"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}

	fmt.Println(resp.GetReply())
}
