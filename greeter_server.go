//go:generate protoc -I ../message --go_out=plugins=grpc:../message ../message/message.proto

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./message"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement message.GreeterServer.
type server struct{}

// Ping implements message.GreeterServer
func (s *server) Ping(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	return &pb.MessageReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
