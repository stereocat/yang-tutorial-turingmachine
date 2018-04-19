package main

import (
	pb "./proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	TapeContent string
}

func (svr *server) Initialize(ctx context.Context, req *pb.InitializeRequest) (*pb.Empty, error) {
	svr.TapeContent = req.GetTapeContent()
	log.Printf("Initialize: TapeContent: %s\n", svr.TapeContent)
	return &pb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	svr := grpc.NewServer()
	pb.RegisterTuringMachineServiceServer(svr, &server{TapeContent: ""})
	reflection.Register(svr)
	if err := svr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
