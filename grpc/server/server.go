package main

import (
	"context"
	"fmt"
	"github.com/yipeng-git/algorithms/grpc/pb"
)

type testServer struct {
	pb.UnimplementedTestServer
}

func newServer() pb.TestServer {
	return &testServer{}
}

func (s *testServer) Echo(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("receive request " + req.String())
	return &pb.EchoResponse{Text: req.Text, Text2: "hello" + req.Text, Num: 888}, nil
}
