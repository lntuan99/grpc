package main

import (
	"fmt"
	"github.com/lnt/grpc/2/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type Server struct {
}

func (server *Server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Printf("Received message body from client: %s", req.GetGreeting())
	return &hello.HelloResponse{Reply: "Reply from the Server!"}, nil
}

func (server *Server) BidiHello(stream hello.HelloService_BidiHelloServer) error {
	for {
		reqs, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			log.Fatalf("Received stream request faild: %v", err)
			return err
		}

		fmt.Println("Greeting from client: ", reqs.GetGreeting())

		res := &hello.HelloResponse{Reply: "Reply from server " + time.Now().Format("15:04:05 02/01/2006")}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8002")

	if err != nil {
		log.Fatalf("Failed to listen on port 8002: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	hello.RegisterHelloServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8002: %v", err)
	}
}
