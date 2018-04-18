package turingmachine

import (
	"encoding/xml"
	"fmt"
)

type RpcInit struct {
	XMLName    xml.Name   `xml:"rpc"`
	Xmlns      string     `xml:"xmlns,attr"`
	MessageId  int        `xml:"message-id,attr"`
	Initialize Initialize `xml:"initialize"`
}
type Initialize struct {
	XMLName     xml.Name `xml:"initialize"`
	Xmlns       string   `xml:"xmlns,attr"`
	TapeContent string   `xml:"tape-content"`
}

var rpcInitString string
var rpcInitStruct *RpcInit

// construct turing machine state
func ReadRpcInitFromFile(xmlFileName string) {
	ReadRpcInitFromString(readXmlString(xmlFileName))
}

// construct turing machine state
func ReadRpcInitFromString(xmlString string) {
	rpcInitString = xmlString
	rpcInitStruct = newRpcInit()
	if verbose {
		rpcInitStruct.PrintXml()
	}
	TMState = newTuringMachineState()
	// change operation state
	doneTapeInitialize()
}

func newRpcInit() *RpcInit {
	var ris = new(RpcInit)
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(rpcInitString), ris); err != nil {
		fmt.Println("!! Error: RPC initialize XML Unmarshal error: ", err)
	}
	return ris
}

func (ris *RpcInit) PrintXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(ris, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}
