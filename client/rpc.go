package tmclient

import (
	pb "../proto"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
)

// SendInit sends InitializeRequest message to Server
func SendInit(ctx context.Context, client pb.TuringMachineRpcClient, initFileName string) {
	log.Printf("Initialize")
	var initRequest *pb.InitializeRequest
	if initFileName != "" {
		initRequest = ReadInitRequestFromFile(initFileName)
	} else {
		log.Fatalf("file: %s not found.", initFileName)
		os.Exit(1)
	}
	log.Printf("tape content: %s\n", initRequest.GetTapeContent())

	if _, err := client.Initialize(ctx, initRequest); err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End initialize\n")
}

// SendConfig sends Configure message to Server
func SendConfig(ctx context.Context, client pb.TuringMachineRpcClient, ttfFileName string) {
	log.Printf("Configure")
	var ttf *pb.Config
	if ttfFileName != "" {
		ttf = ReadTtfFromFile(ttfFileName)
	} else {
		log.Fatalf("file: %s not found.", ttfFileName)
		os.Exit(1)
	}

	if _, err := client.Configure(ctx, ttf); err != nil {
		log.Fatalf("could not configure: %v\n", err)
	}
	log.Printf("End configure\n")
}

// SendRun sends Run message to Server
func SendRun(ctx context.Context, client pb.TuringMachineRpcClient) {
	log.Printf("Run")
	halted, err := client.Run(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run, state=%d\n", halted.GetState())
}

// SendGetState sends GetState message to Server
func SendGetState(ctx context.Context, client pb.TuringMachineRpcClient) {
	log.Printf("Get State of Server Turing Machine\n")
	tm, err := client.GetState(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get state: %v\n", err)
	}
	fmt.Println(tm.ToXmlString())
	log.Printf("End Get State\n")
}
