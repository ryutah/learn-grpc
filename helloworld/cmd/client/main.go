package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ryutah/learn-grpc/helloworld/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "ryutah",
	})
	if err != nil {
		panic(err)
	}

	payload, _ := json.MarshalIndent(reply, "", "  ")
	fmt.Println(string(payload))
}
