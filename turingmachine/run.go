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
var rpcInitStruct *Rpc
var TMState *TuringMachineState

func ReadRpcInitFromFile(xmlFileName string) {
	// construct turing machine state
	rpcInitString = readXmlString(xmlFileName)
	rpcInitStruct = newRpc()
	TMState = newTuringMachineState()
}

func newRpc() *Rpc {
    var ris = new(Rpc)
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(rpcInitString), ris); err != nil {
		fmt.Println("!! Error: RPC initialize XML Unmarshal error: ", err)
	}
    return ris
}

func (ris *Rpc) PrintXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(ris, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}

func (tmState *TuringMachineState) toString(step int) string {
	var stateString string
	stateString = fmt.Sprintf("%4d  [S%d] | ", step, tmState.State)
	for i, cell := range tmState.Tape.CellList {
		if i == int(tmState.HeadPosition) {
			stateString += fmt.Sprintf("<%s>|", cell.Symbol)
		} else {
			stateString += fmt.Sprintf(" %s |", cell.Symbol)
		}
	}
	return stateString
}

func (tmState *TuringMachineState) setState(action Output) {
	// change to next state
	tmState.State = action.State
	// write symbol to tape under head-position
	if action.Symbol != "" {
		tmState.Tape.CellList[tmState.HeadPosition].Symbol = action.Symbol
	}
}

func (tmState *TuringMachineState) setHeadPosition(action Output) {
	// move to next head position
	switch action.HeadMove {
	case "left":
		tmState.HeadPosition -= 1
	case "right":
		tmState.HeadPosition += 1
	}
}

func (action Output) toString() string {
	var moveString string
	switch action.HeadMove {
	case "left":
		moveString = "<= "
	case "right":
		moveString = " =>"
	}
	return fmt.Sprintf(" [S%d] %5s %4s", action.State, action.Symbol, moveString)
}

func (tmState *TuringMachineState) Run() {
	var step = 1
	var finishState = TransitionTable.GetFinishState()

	// header
	var indent = len([]byte(tmState.toString(step))) - 20 // offset
	var headerString = fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	for tmState.State != finishState {
		var currentString = tmState.toString(step)
		// read current symbol under head-position
		var currentSymbol = tmState.Tape.CellList[tmState.HeadPosition].Symbol
		// find next action by current state and symbol
		var action = TransitionTable[tmState.State][currentSymbol]
		// print step
		fmt.Println(currentString + action.toString())
		// Go to Next: change state and head-position
		tmState.setState(action)
		tmState.setHeadPosition(action)
		step++
	}
	fmt.Println(tmState.toString(step) + " END")
}

func newTuringMachineState() *TuringMachineState {
    var tmState = new(TuringMachineState)
	var content = []byte(rpcInitStruct.Initialize.TapeContent) // string 2 []byte
	tmState.State = 0
	tmState.HeadPosition = 1
	var cellList = make([]Cell, 0)
	for coord, byteSymbol := range content {
		var cell = Cell{Coord: coord, Symbol: string(byteSymbol)}
		cellList = append(cellList, cell)
	}
	tmState.Tape = Tape{CellList: cellList}
	// *tmState.TransitionFunction = TransitionTable // TBA
    return tmState
}

func (tmState *TuringMachineState) PrintXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tmState, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}
