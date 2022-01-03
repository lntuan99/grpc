package main

import (
	"fmt"
	"github.com/lnt/grpc/2/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	stream, err := client.BidiHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var count = 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			if count == 0 {
				stream.CloseSend()
				break
			}

			count -= 1

			req := hello.HelloRequest{Greeting: fmt.Sprintf("%v_%v", time.Now().Unix(), rand.Int())}
			if err := stream.Send(&req); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			fmt.Println(res.GetReply())
		}
	}()

	wg.Wait()
}
