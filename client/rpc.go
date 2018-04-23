package tm_client

import (
	pb "../proto"
	"fmt"
	context "golang.org/x/net/context"
	"log"
	"os"
)

func SendInit(ctx context.Context, client pb.TuringMachineRpcClient, initFileName string) {
	log.Printf("Initialize")
	var initRequest *pb.InitializeRequest
	if initFileName != "" {
		initRequest = ReadInitRequestFromFile(initFileName)
	} else {
		log.Fatalf("file: %s not found.", initFileName)
		os.Exit(1)
	}
	log.Printf("# initReq: %v\n", initRequest)

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
	log.Printf("Run")
	halted, err := client.Run(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run, state=%d\n", halted.GetState())
}

func SendGetState(ctx context.Context, client pb.TuringMachineRpcClient) {
	log.Printf("Get State of Server Turing Machine\n")
	tm, err := client.GetState(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get state: %v\n", err)
	}
	fmt.Println(TMXmlString(tm))
	log.Printf("End Get State\n")
}
