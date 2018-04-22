package tm_client

import (
	pb "../proto"
	context "golang.org/x/net/context"
	"log"
	"os"
)

func SendInit(ctx context.Context, client pb.TuringMachineRpcClient, tape string) {
	// Initialize
	log.Printf("Initialize")
	initRequest := &pb.InitializeRequest{TapeContent: tape}
	_, err := client.Initialize(ctx, initRequest) // return Empty
	if err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End initialize\n")
}

func SendConfig(ctx context.Context, client pb.TuringMachineRpcClient, ttfFileName string) {
	log.Printf("Configure")
	var ttf *pb.Config
	if ttfFileName != "" {
		ttf = ReadTtfFromFile(ttfFileName)
	} else {
		log.Fatalf("file: %s not found.", ttfFileName)
		os.Exit(1)
	}
	_, err := client.Configure(ctx, ttf)
	if err != nil {
		log.Fatalf("could not configure: %v\n", err)
	}
	log.Printf("End configure\n")
}

func SendRun(ctx context.Context, client pb.TuringMachineRpcClient) {
	// Run
	log.Printf("Run")
	_, err := client.Run(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run")
}
