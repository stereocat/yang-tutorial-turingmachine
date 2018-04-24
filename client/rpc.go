package tmclient

import (
	pb "../proto"
	"fmt"
	"log"
)

// SendInit sends InitializeRequest message to Server
func (tmClient *TMClient) SendInit() {
	log.Printf("Initialize")
	var initRequest *pb.InitializeRequest
	if tmClient.InitFileName != "" {
		initRequest = ReadInitRequestFromFile(tmClient.InitFileName)
	} else {
		log.Fatalf("file: %s not found.", tmClient.InitFileName)
	}
	log.Printf("tape content: %s\n", initRequest.GetTapeContent())

	if _, err := tmClient.Client.Initialize(tmClient.Ctx, initRequest); err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End initialize\n")
}

// SendConfig sends Configure message to Server
func (tmClient *TMClient) SendConfig() {
	log.Printf("Configure")
	var ttf *pb.Config
	if tmClient.TtfFileName != "" {
		ttf = ReadTtfFromFile(tmClient.TtfFileName)
	} else {
		log.Fatalf("file: %s not found.", tmClient.TtfFileName)
	}

	if _, err := tmClient.Client.Configure(tmClient.Ctx, ttf); err != nil {
		log.Fatalf("could not configure: %v\n", err)
	}
	log.Printf("End configure\n")
}

// SendRun sends Run message to Server
func (tmClient *TMClient) SendRun() {
	log.Printf("Run")
	halted, err := tmClient.Client.Run(tmClient.Ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run, state=%d\n", halted.GetState())
}

// SendGetState sends GetState message to Server
func (tmClient *TMClient) SendGetState() {
	log.Printf("Get State of Server Turing Machine\n")
	tm, err := tmClient.Client.GetState(tmClient.Ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get state: %v\n", err)
	}
	fmt.Println(tm.ToXmlString())
	log.Printf("End Get State\n")
}
