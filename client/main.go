package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "example.com/helloRPC/chat"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewMyFirstServiceClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	n, err2 := c.Other(ctx, &pb.DataReq{Amount: 21})
	if err2 != nil {
		log.Fatalf("could not greet: %v", n)
	}
	log.Printf("Amount: %d", n.Amount)

	// r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())

}
