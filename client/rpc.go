package tmclient

import (
	pb "../proto"
	"fmt"
	"log"
)

// SendInit sends InitializeRequest message to Server
func (tmClient *TMClient) SendInit(initRequest *pb.InitializeRequest) {
	log.Printf("initialize\n")
	if _, err := tmClient.Client.Initialize(tmClient.Ctx, initRequest); err != nil {
		log.Fatalf("could not initialize: %v\n", err)
	}
	log.Printf("End initialize\n")
}

// SendConfig sends Configure message to Server
func (tmClient *TMClient) SendConfig(config *pb.Config) {
	log.Printf("Configure")
	if _, err := tmClient.Client.Configure(tmClient.Ctx, config); err != nil {
		log.Fatalf("could not configure: %v\n", err)
	}
	log.Printf("End configure\n")
}

// SendRun sends Run message to Server
func (tmClient *TMClient) SendRun() {
	log.Printf("Run")
	notification, err := tmClient.Client.Run(tmClient.Ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not run: %v\n", err)
	}
	log.Printf("End Run, machine state=%d\n", notification.GetHalted().GetState())
	fmt.Println(notification.ToXMLString())
}

// SendGetState sends GetState message to Server
func (tmClient *TMClient) SendGetState() {
	log.Printf("Get State of Server Turing Machine\n")
	tm, err := tmClient.Client.GetState(tmClient.Ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get state: %v\n", err)
	}
	fmt.Println(tm.ToXMLString())
	log.Printf("End Get State\n")
}
