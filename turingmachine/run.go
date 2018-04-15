package turingmachine

import (
	"encoding/xml"
	"fmt"
)

type Rpc struct {
	XMLName    xml.Name   `xml:"rpc"`
	Initialize Initialize `xml:"initialize"`
}
type Initialize struct {
	XMLName     xml.Name `xml:"initialize"`
	TapeContent string   `xml:"tape-content"`
}

func ParseRpcInitializeString(xmlString string) Rpc {
	var data = Rpc{}
	if err := xml.Unmarshal([]byte(xmlString), &data); err != nil {
		fmt.Println("!! RPC initialize XML Unmarshal error: ", err)
	}
	return data
}

func PrintRpcInitXmlByStruct(xmlStruct Rpc) {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(xmlStruct, "", "  ")
	if err != nil {
		fmt.Println("!! XML Marshal err: ", err)
	}
	fmt.Println("## xml data")
	fmt.Println(string(xmlBuf))
}

type TuringMachineState struct {
	XMLName            xml.Name           `xml:"turing-machine"`
	State              uint16             `xml:"state"`
	HeadPosition       uint16             `xml:"head-position"`
	Tape               Tape               `xml:"tape"`
	TransitionFunction TransitionFunction `xml:"transition-function"`
}
type Tape struct {
	CellList []Cell `xml:"cell"`
}
type Cell struct {
	Coord  int    `xml:"coord"`
	Symbol string `xml:"symbol"`
}

var TMState TuringMachineState

func RunTuringMachine() {
	var step = 1
	var finishState = GetFinishState()

	for TMState.State != finishState {
		// read current symbol under head-position
		var currentHeadPosition = TMState.HeadPosition
		var currentSymbol = TMState.Tape.CellList[currentHeadPosition].Symbol
		var currentState = TMState.State
		// find next action by current state and symbol
		var action = TransitionTable[currentState][currentSymbol]

		// change to next state
		TMState.State = action.State
		// write symbol to tape under head-position
		if action.Symbol != "" {
			TMState.Tape.CellList[currentHeadPosition].Symbol = action.Symbol
		}
		// move to next head position
		switch action.HeadMove {
		case "left":
			TMState.HeadPosition -= 1
		case "right":
			TMState.HeadPosition += 1
		}

		fmt.Println("step: ", step)
		fmt.Printf("  current: head: %d, state:%d, symbol:%s\n", currentHeadPosition, currentState, currentSymbol)
		fmt.Printf("  action : head: %d, state:%d, symbol:%s, move:%s\n", TMState.HeadPosition, action.State, action.Symbol, action.HeadMove)

		step++
	}
}

func InitializeTuringMachine(iniXmlStruct Rpc) {
	var content = []byte(iniXmlStruct.Initialize.TapeContent) // string 2 []byte
	TMState.State = 0
	TMState.HeadPosition = 1 // FIX later
	var cellList = make([]Cell, 0)
	for coord, byteSymbol := range content {
		var cell = Cell{Coord: coord, Symbol: string(byteSymbol)}
		cellList = append(cellList, cell)
	}
	TMState.Tape = Tape{CellList: cellList}
	// TMState.TransitionFunction = TransitionTable // TBA
}

func PrintTMStateXmlByStruct(xmlStruct TuringMachineState) {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(xmlStruct, "", "  ")
	if err != nil {
		fmt.Println("!! XML Marshal err: ", err)
	}
	fmt.Println("## xml data")
	fmt.Println(string(xmlBuf))
}
