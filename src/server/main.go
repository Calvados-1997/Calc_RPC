package main

import (
	"context"
	"fmt"
	"log"
	"net"

	hellopub "github.com/Calvados-1997/Calc_RPC/pkg/gRPC"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	hellopub.RegisterGreetingServiceServer(s, NewMyServer())

	reflection.Register(s)

	log.Printf("start gRPC server port on %d", port)
	s.Serve(listener)
}

type myServer struct {
	hellopub.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopub.HelloRequest) (*hellopub.HelloResponse, error) {
	return &hellopub.HelloResponse{
		Message: fmt.Sprintf("Hello, %s from server!", req.GetName()),
	}, nil
}
