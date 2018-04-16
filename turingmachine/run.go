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

var rpcInitString string
var rpcInitStruct Rpc
var TMState TuringMachineState

func ReadRpcInitFromFile(xmlFileName string) {
	// construct turing machine state
	rpcInitString = readXmlString(xmlFileName)
	parseRpcInitString()
	initializeTuringMachine()
}

func parseRpcInitString() {
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(rpcInitString), &rpcInitStruct); err != nil {
		fmt.Println("!! Error: RPC initialize XML Unmarshal error: ", err)
	}
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

func getTuringMachineString(step int) string {
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

func setNextState(action Output) {
	// change to next state
	TMState.State = action.State
	// write symbol to tape under head-position
	if action.Symbol != "" {
		TMState.Tape.CellList[TMState.HeadPosition].Symbol = action.Symbol
	}
}

func setNextHeadPosition(action Output) {
	// move to next head position
	switch action.HeadMove {
	case "left":
		TMState.HeadPosition -= 1
	case "right":
		TMState.HeadPosition += 1
	}
}

func getNextActionString(action Output) string {
	var moveString string
	switch action.HeadMove {
	case "left":
		moveString = "<= "
	case "right":
		moveString = " =>"
	}
	return fmt.Sprintf(" [S%d] %5s %4s", action.State, action.Symbol, moveString)
}

func RunTuringMachine() {
	var step = 1
	var finishState = GetFinishState()

	// header
	var indent = len([]byte(getTuringMachineString(step))) - 20 // offset
	var headerString = fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	for TMState.State != finishState {
		var currentString = getTuringMachineString(step)
		// read current symbol under head-position
		var currentSymbol = TMState.Tape.CellList[TMState.HeadPosition].Symbol
		// find next action by current state and symbol
		var action = TransitionTable[TMState.State][currentSymbol]
		// change state and head-position
		setNextState(action)
		setNextHeadPosition(action)
		// print step
		fmt.Println(currentString + getNextActionString(action))

		step++
	}
	fmt.Println(getTuringMachineString(step) + " END")
}

func initializeTuringMachine() {
	var content = []byte(rpcInitStruct.Initialize.TapeContent) // string 2 []byte
	TMState.State = 0
	TMState.HeadPosition = 1
	var cellList = make([]Cell, 0)
	for coord, byteSymbol := range content {
		var cell = Cell{Coord: coord, Symbol: string(byteSymbol)}
		cellList = append(cellList, cell)
	}
	TMState.Tape = Tape{CellList: cellList}
	// TMState.TransitionFunction = TransitionTable // TBA
}

func PrintTMStateXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(TMState, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println("## xml data")
	fmt.Println(string(xmlBuf))
}
