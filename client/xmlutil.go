package tmclient

import (
	pb "../proto"
	"bufio"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

// common utilities

func stringFromDataFile(fileName string) string {
	filePtr, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error: Cannot read file:%s\n", fileName)
	} else {
		defer filePtr.Close() // close finished readXmlFile
	}
	return openDataFile(filePtr)
}

func openDataFile(filePtr *os.File) string {
	// read data from filePtr
	var scanner = bufio.NewScanner(filePtr)
	var lines = make([]string, 0) // multiple lines
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n") // convert to single line
}

// for initial request

// ReadInitRequestFromFile reads transition table function data from file
func (tmClient *TMClient) ReadInitRequestFromFile(fileName string) *pb.InitializeRequest {
	if _, err := os.Stat(fileName); err == nil {
		// if directly specified filename at command
		return tmClient.NewInitRequest(stringFromDataFile(fileName))
	} else if tmClient.InitFileName != "" {
		// if filename not specified but specify default by client argv
		return tmClient.ReadInitRequestFromFile(tmClient.InitFileName)
	}
	// if can not clear filename, use stdin
	return tmClient.NewInitRequest(readXMLStringFromStdin())
}

// NewInitRequest returns RPC InitializeRequest message
func (tmClient *TMClient) NewInitRequest(dataString string) *pb.InitializeRequest {
	rpcInitReq := new(pb.Rpc) // initialize request struct
	if tmClient.UseJSON {
		if err := json.Unmarshal([]byte(dataString), rpcInitReq); err != nil {
			log.Printf("Error: Initialize Request JSON Unmarshal error: %v\n", err)
		}
	} else {
		if err := xml.Unmarshal([]byte(dataString), rpcInitReq); err != nil {
			log.Printf("Error: Initialize Request XML Unmarshal error: %v\n", err)
		}
	}
	return rpcInitReq.GetInitialize()
}

// for TTF(TuringMachine)

// ReadTuringMachineFromFile reads Transition Table data from file
// to configure Turing Machine (constructor)
func (tmClient *TMClient) ReadTuringMachineFromFile(fileName string) *pb.TuringMachine {
	if _, err := os.Stat(fileName); err == nil {
		// if directly specified filename at command
		return tmClient.NewTuringMachine(stringFromDataFile(fileName))
	} else if tmClient.TtfFileName != "" {
		// if filename not specified but specify default by client argv
		return tmClient.NewTuringMachine(stringFromDataFile(tmClient.TtfFileName))
	}
	// if can not clear filename, use stdin
	return tmClient.NewTuringMachine(readXMLStringFromStdin())
}

// NewTuringMachine returns Transition Table data from XML string
func (tmClient *TMClient) NewTuringMachine(dataString string) *pb.TuringMachine {
	config := new(pb.Config) // transition table struct
	if tmClient.UseJSON {
		if err := json.Unmarshal([]byte(dataString), config); err != nil {
			log.Printf("Error: TransitionTable JSON Unmarshal error: %v\n", err)
		}
	} else {
		if err := xml.Unmarshal([]byte(dataString), config); err != nil {
			log.Printf("Error: TransitionTable XML Unmarshal error: %v\n", err)
		}
	}
	return config.GetTuringMachine()
}
