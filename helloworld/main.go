package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ryutah/learn-grpc/helloworld/helloworld"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-sigCtx.Done()
		log.Println("shutdown server...")
		s.GracefulStop()
	}()

	helloworld.RegisterGreeterServer(s, new(server))
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to start server: %v", err)
	}
}

type server struct {
	helloworld.UnimplementedGreeterServer
}

var _ helloworld.GreeterServer = new(server)

// Sends a greeting
func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("got request: %s ", mustMarshalJSON(req))
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("Hello %s !!", req.GetName()),
	}, nil
}

func mustMarshalJSON(v interface{}) []byte {
	payload, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return payload
}
