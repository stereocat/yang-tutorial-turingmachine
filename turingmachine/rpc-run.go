package turingmachine

import (
	"encoding/xml"
	"fmt"
)

type TuringMachineState struct {
	XMLName            xml.Name           `xml:"turing-machine"`
	Xmlns              string             `xml:"xmlns,attr"`
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

var TMState *TuringMachineState

func (tmState *TuringMachineState) Print() {
	fmt.Println(tmState.toString(0))
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

func (tmState *TuringMachineState) Run() *Notification {
	var (
		step        = 1
		finishState = TransitionTable.GetFinishState()
		// header
		indent       = len([]byte(tmState.toString(step))) - 20 // offset
		headerString = fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	)
	fmt.Printf(headerString, " ")

	// Run
	for tmState.State != finishState {
		var (
			// current Turing Machine state
			currentString = tmState.toString(step)
			// read symbol under head-position
			currentSymbol = tmState.Tape.CellList[tmState.HeadPosition].Symbol
			// find next action by current state and symbol
			action = TransitionTable[tmState.State][currentSymbol]
		)
		// print step
		fmt.Println(currentString + action.toString())
		// Go to Next: change state and head-position
		tmState.setState(action)
		tmState.setHeadPosition(action)
		step++
	}
	fmt.Println(tmState.toString(step) + " END")

	return newNotification(tmState.State)
}

func newTuringMachineState() *TuringMachineState {
	var (
		content  = []byte(rpcInitStruct.Initialize.TapeContent) // string 2 []byte
		cellList = make([]Cell, 0)
	)
	for coord, byteSymbol := range content {
		var cell = Cell{Coord: coord, Symbol: string(byteSymbol)}
		cellList = append(cellList, cell)
	}
	return &TuringMachineState{
		State:              0,
		HeadPosition:       1,
		Xmlns:              "http://example.net/turing-machine",
		Tape:               Tape{CellList: cellList},
		TransitionFunction: transitionTableStruct.TuringMachine.TransitionFunction,
	}
}

func (tmState *TuringMachineState) PrintXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tmState, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}
