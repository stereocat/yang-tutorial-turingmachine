package main

import (
	pb "./proto"
	tms "./server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	svr := grpc.NewServer()
	// Register initial state of TMServer struct into RpcServer
	pb.RegisterTuringMachineRpcServer(svr, tms.NewTMServer())
	reflection.Register(svr)
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
