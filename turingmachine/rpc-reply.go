package turingmachine

import (
	"encoding/xml"
	"fmt"
)

type RpcReply struct {
	XMLName   xml.Name     `xml:"rpc-reply"`
	Xmlns     string       `xml:"xmlns,attr"`
	MessageId int          `xml:"message-id,attr"`
	Data      RpcReplyData `xml:"data"`
}

type RpcReplyData struct {
	XMLName       xml.Name           `xml:"data"`
	TuringMachine TuringMachineState `xml:"turing-machine"`
}

func GetConfig() *RpcReply {
	return newRpcReply()
}

func newRpcReply() *RpcReply {
	messageId += 1
	return &RpcReply{
		Xmlns:     "http://example.net/turing-machine",
		MessageId: messageId,
		Data: RpcReplyData{
			TuringMachine: *TMState,
		},
	}
}

func (rpcRep *RpcReply) PrintXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(rpcRep, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}
