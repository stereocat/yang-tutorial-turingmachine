package tmclient

import (
	pb "../proto"
	"bufio"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

// common utilities

func stringFromXMLFile(xmlFileName string) string {
	xmlFile, err := os.Open(xmlFileName)
	if err != nil {
		log.Fatalf("Error: Cannot read file:%s\n", xmlFileName)
	} else {
		defer xmlFile.Close() // close finished readXmlFile
	}
	return openXMLFile(xmlFile)
}

func openXMLFile(xmlFile *os.File) string {
	// read data from file
	var scanner = bufio.NewScanner(xmlFile)
	var lines = make([]string, 0) // multiple lines
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n") // convert to single line
}

// for initial request

// ReadInitRequestFromFile reads transition table function data from file
func ReadInitRequestFromFile(xmlFileName string) *pb.InitializeRequest {
	return NewInitRequest(stringFromXMLFile(xmlFileName))
}

// NewInitRequest returns RPC InitializeRequest message
func NewInitRequest(xmlString string) *pb.InitializeRequest {
	rpcInitReq := new(pb.Rpc) // initialize request struct
	// unmarshal (parse): xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), rpcInitReq); err != nil {
		log.Printf("Error: Initialize Request XML Unmarshal error: %v\n", err)
	}
	return rpcInitReq.GetInitialize()
}

// for TTF(Config)

// ReadTtfFromFile reads Transition Table data from file
// to configure Turing Machine (constructor)
func ReadTtfFromFile(xmlFileName string) *pb.Config {
	return NewConfig(stringFromXMLFile(xmlFileName))
}

// NewConfig returns Transition Table data from XML string
func NewConfig(xmlString string) *pb.Config {
	config := new(pb.Config) // transition table struct
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), config); err != nil {
		log.Printf("Error: TransitionTable XML Unmarshal error: %v\n", err)
	}
	return config
}
