package main

import (
	"github.com/lnt/grpc/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("Failed to listen on port 8001: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8001: %v", err)
	}

}
