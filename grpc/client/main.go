package main

import (
	"context"
	"fmt"
	"github.com/yipeng-git/algorithms/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

var client pb.TestClient
var addr = ":50001"

func main() {
	_ = map[string]string{}
	initClient()

	res, err := client.Echo(context.Background(), &pb.EchoRequest{Text: "hello"})
	if err != nil {
		panic("echo err " + err.Error())
	}
	fmt.Println(res)
}

func initClient() {
	param := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
	option := grpc.WithKeepaliveParams(param)
	//fmt.Println("before dial")
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), option)
	//fmt.Println("after dial")
	if err != nil {
		fmt.Printf("grpc dial error %v\n", err)
	}
	client = pb.NewTestClient(conn)
}
