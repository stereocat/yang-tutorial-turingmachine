package tmclient

import (
	pb "../proto"
	"encoding/xml"
	"log"
)

// ReadInitRequestFromFile reads transition table function data from file
func ReadInitRequestFromFile(xmlFileName string) *pb.InitializeRequest {
	return readInitRequestFromString(stringFromXMLFile(xmlFileName))
}

func readInitRequestFromString(xmlString string) *pb.InitializeRequest {
	return NewInitRequest(xmlString)
}

// NewInitRequest returns RPC InitializeRequest message
func NewInitRequest(xmlString string) *pb.InitializeRequest {
	rpcInitReq := new(pb.Rpc) // initialize request struct
	// unmarshal (parse): xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), rpcInitReq); err != nil {
		log.Fatalf("Error: Initialize Request XML Unmarshal error: %v\n", err)
	}
	return rpcInitReq.GetInitialize()
}
