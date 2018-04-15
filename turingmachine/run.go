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
		fmt.Println("!! Error: RPC initialize XML Unmarshal error: ", err)
	}
	return data
}

func PrintRpcInitXmlByStruct(xmlStruct Rpc) {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(xmlStruct, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
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

func GetTuringMachineString(step int) string {
	var stateString string
	stateString = fmt.Sprintf("%4d  [S%d] | ", step, TMState.State)
	for i, cell := range TMState.Tape.CellList {
		if i == int(TMState.HeadPosition) {
			stateString += fmt.Sprintf("<%s>|", cell.Symbol)
		} else {
			stateString += fmt.Sprintf(" %s |", cell.Symbol)
		}
	}
	return stateString
}

func RunTuringMachine() {
	var step = 1
	var finishState = GetFinishState()

	// header
	var indent = len([]byte(GetTuringMachineString(step))) - 20
	var headerString = fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	for TMState.State != finishState {
		var currentString = GetTuringMachineString(step)

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

		var moveString string
		switch action.HeadMove {
		case "left":
			moveString = "<= "
		case "right":
			moveString = " =>"
		}
		var nextString = fmt.Sprintf(" [S%d] %5s %4s", action.State, action.Symbol, moveString)
		fmt.Println(currentString + nextString)

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
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println("## xml data")
	fmt.Println(string(xmlBuf))
}
