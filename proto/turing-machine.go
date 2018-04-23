package turing_machine

import (
	"encoding/xml"
	"fmt"
	"log"
)

func (output *TuringMachine_TransitionFunction_Delta_Output) ToString() string {
	var moveString string
	switch output.GetHeadMove() {
	case "left":
		moveString = "<= "
	case "right":
		moveString = " =>"
	}
	return fmt.Sprintf(" [S%d] %5s %4s", output.GetState(), output.GetSymbol(), moveString)
}

func (tm *TuringMachine) ToString(step int) string {
	var stateString string
	stateString = fmt.Sprintf("%4d  [S%d] | ", step, tm.GetState())
	for i, cell := range tm.GetTape().GetCell() {
		if i == int(tm.GetHeadPosition()) {
			stateString += fmt.Sprintf("<%s>|", cell.GetSymbol())
		} else {
			stateString += fmt.Sprintf(" %s |", cell.GetSymbol())
		}
	}
	return stateString
}

func (tm *TuringMachine) ChangeState(action *TuringMachine_TransitionFunction_Delta_Output) {
	// change to next state
	tm.State = action.GetState()
	// write symbol to tape under head-position
	if action.GetSymbol() != "" {
		cellList := tm.GetTape().GetCell()
		cellList[tm.GetHeadPosition()].Symbol = action.GetSymbol()
	}
	// move to next head position
	switch action.GetHeadMove() {
	case "left":
		tm.HeadPosition -= 1
	case "right":
		tm.HeadPosition += 1
	}
}

// TMXmlString
func (tm *TuringMachine) ToXmlString() string {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tm, "", "  ")
	if err != nil {
		log.Fatalf("Error: XML Marshal err: ", err)
	}
	return string(xmlBuf)
}
