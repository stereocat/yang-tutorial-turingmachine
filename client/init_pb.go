package tm_client

import (
	pb "../proto"
	"encoding/xml"
	"fmt"
	"log"
)

func ReadInitRequestFromFile(xmlFileName string) *pb.InitializeRequest {
	return readInitRequestFromString(stringFromXmlFile(xmlFileName))
}

func readInitRequestFromString(xmlString string) *pb.InitializeRequest {
	return NewInitRequest(xmlString)
}

func NewInitRequest(xmlString string) *pb.InitializeRequest {
	rpcInitReq := new(pb.Rpc) // initialize request struct
	fmt.Printf("# xmlString: %s\n", xmlString)
	// unmarshal (parse): xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), rpcInitReq); err != nil {
		log.Fatalf("Error: Initialize Request XML Unmarshal error: %v\n", err)
	}
	return rpcInitReq.GetInitialize()
}
