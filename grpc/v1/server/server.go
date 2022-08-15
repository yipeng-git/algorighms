package main

import (
	"context"
	"github.com/yipeng-git/algorithms/grpc/v1/pb"
)

type testServer struct {
	pb.UnimplementedTestServer
}

func newServer() pb.TestServer {
	return &testServer{}
}

func (s *testServer) Echo(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Text: req.Text}, nil
}
