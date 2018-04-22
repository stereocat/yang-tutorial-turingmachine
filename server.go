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
	// Regist initial state of Server struct
	// to avoid empty(nil) function calling
	pb.RegisterTuringMachineRpcServer(svr, &tms.Server{
		TuringMachine: &pb.TuringMachine{
			HeadPosition: 1,
			State:        0,
			Tape: &pb.TuringMachine_Tape{
				Cell: make([]*pb.TuringMachine_Tape_Cell, 0),
			},
			TransitionFunction: &pb.TuringMachine_TransitionFunction{},
		},
		TransitionTable: tms.TTF{}, // map(ref-type)
	})
	reflection.Register(svr)
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
