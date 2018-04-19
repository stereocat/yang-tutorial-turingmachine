package main

import (
	pb "./proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewTuringMachineServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	initRequest := &pb.InitializeRequest{TapeContent: "0111011"}
	_, err = client.Initialize(ctx, initRequest) // return Empty
	log.Printf("Initialize")
	if err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End\n")
}
