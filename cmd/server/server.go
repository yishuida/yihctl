package main

import (
	log "github.com/sirupsen/logrus"
	myrpc "github.com/yishuida/yihctl/cmd/rpc"
	"google.golang.org/grpc"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	myrpc.RegisterHelloServiceServer(grpcServer, new(myrpc.HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}