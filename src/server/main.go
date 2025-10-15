package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "github.com/Calvados-1997/Calc_RPC/pkg/gRPC/helloapp/v1"
	"github.com/Calvados-1997/Calc_RPC/pkg/gRPC/helloapp/v1/gRPCconnect"
)

type GreetServer struct{}

func (s *GreetServer) Hello(
	ctx context.Context,
	req *connect.Request[greetv1.HelloRequest],
) (*connect.Response[greetv1.HelloResponse], error) {
	res := connect.NewResponse(&greetv1.HelloResponse{
		Message: fmt.Sprintf("Hello %s from Server!", req.Msg.Name),
	})
	return res, nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, connect-protocol-version")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := gRPCconnect.NewGreetingServiceHandler(greeter)
	mux.Handle(path, corsMiddleware(handler))
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
