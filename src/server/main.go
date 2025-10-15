package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	gRPC "github.com/Calvados-1997/Calc_RPC/pkg/gRPC/helloapp/v1"
	greetv1 "github.com/Calvados-1997/Calc_RPC/pkg/gRPC/helloapp/v1"
	"github.com/Calvados-1997/Calc_RPC/pkg/gRPC/helloapp/v1/gRPCconnect"
)

type GreetServer struct{}

func (s *GreetServer) Hello(
	ctx context.Context,
	req *connect.Request[gRPC.HelloRequest],
) (*connect.Response[gRPC.HelloResponse], error) {
	res := connect.NewResponse(&greetv1.HelloResponse{
		Message: fmt.Sprintf("Hello %s from Server!", req.Msg.Name),
	})
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := gRPCconnect.NewGreetingServiceHandler(greeter)
	mux.Handle(path, handler)
	p := new(http.Protocols)
	p.SetHTTP1(true)

	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}
	fmt.Printf("Start server on %s", s.Addr)
	s.ListenAndServe()
}
