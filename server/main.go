package main

import (
	"context"
	"log"
	"net"

	pb "example.com/helloRPC/chat"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyFirstServiceServer
}

const (
	port = ":50051"
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received (again): %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) Other(ctx context.Context, in *pb.DataReq) (*pb.DataRes, error) {
	log.Printf("Received (other): %v", "ggg")
	return &pb.DataRes{Amount: 4}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMyFirstServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
