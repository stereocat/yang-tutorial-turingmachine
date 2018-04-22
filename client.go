package main

import (
	tmc "./client"
	pb "./proto"
	"flag"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

var (
	ttfFileOpt = flag.String("t", "", "transition table function data xml")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewTuringMachineRpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Config
	log.Printf("Configure")
	var ttf *pb.Config
	if *ttfFileOpt != "" {
		ttf = tmc.ReadTtfFromFile(*ttfFileOpt)
	}
	_, err = client.Configure(ctx, ttf)
	if err != nil {
		log.Fatalf("could not configure: %v\n", err)
	}
	log.Printf("End configure\n")

	// Initialize
	log.Printf("Initialize")
	initRequest := &pb.InitializeRequest{TapeContent: "0111011"}
	_, err = client.Initialize(ctx, initRequest) // return Empty
	if err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End initialize\n")

	// Run
	log.Printf("Run")
	_, err = client.Run(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run")
}
