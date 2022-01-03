package main

import (
	"context"
	"github.com/lnt/grpc/2/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	message := hello.HelloRequest{Greeting: "Hello from client"}

	response, err := client.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.GetReply())
}
