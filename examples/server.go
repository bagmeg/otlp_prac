package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bagmeg/otlp_prac/data"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTraceServer
}

func (s *Server) Consume(_ context.Context, t *pb.TraceData) (*pb.Reply, error) {
	return &pb.Reply{
		Message: "Success",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTraceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
