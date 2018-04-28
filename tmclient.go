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
	ttfFileOpt  = flag.String("t", "", "transition table function data xml")
	initFileOpt = flag.String("i", "", "rpc initialize xml")
	useJson     = flag.Bool("j", false, "Use JSON instead of XML")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewTuringMachineRpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute) // timeout 5 minute
	defer cancel()

	log.Printf("Start CLI")
	tmClient := tmc.NewTMClient(ctx, client, *ttfFileOpt, *initFileOpt, *useJson)
	tmClient.StartCli()
}
