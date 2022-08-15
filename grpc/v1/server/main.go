package main

import (
	"github.com/yipeng-git/algorithms/grpc/v1/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := newServer()
	s := grpc.NewServer()
	pb.RegisterTestServer(s, server)
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen grpc port %v error %v", ":50001", err)
	}
	_ = s.Serve(lis)
}
