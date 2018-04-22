package main

import (
	pb "./proto"
	tms "./server"
	grpc "google.golang.org/grpc"
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
	pb.RegisterTuringMachineRpcServer(svr, &tms.Server{})
	reflection.Register(svr)
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
