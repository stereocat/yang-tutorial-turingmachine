package tm_server

import (
	pb "../proto"
	"fmt"
)

const STEP_MAX = 100

func getTMString(tm *pb.TuringMachine, step int) string {
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

func changeTMState(tm *pb.TuringMachine, action *pb.TuringMachine_TransitionFunction_Delta_Output) {
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

func getOutputString(action *pb.TuringMachine_TransitionFunction_Delta_Output) string {
	var moveString string
	switch action.GetHeadMove() {
	case "left":
		moveString = "<= "
	case "right":
		moveString = " =>"
	}
	return fmt.Sprintf(" [S%d] %5s %4s", action.GetState(), action.GetSymbol(), moveString)
}

func (svr *Server) RunTM() uint32 {
	tm := svr.TuringMachine
	step := 1
	finishState := svr.TransitionTable.GetFinishState()

	// header
	indent := len([]byte(getTMString(tm, step))) - 20 // offset
	headerString := fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	for tm.GetState() != finishState && step < STEP_MAX {
		var (
			// read symbol under head-position
			cellList      = tm.GetTape().GetCell()
			currentSymbol = cellList[tm.GetHeadPosition()].GetSymbol()
			// find next action by current state and symbol
			action = svr.TransitionTable[tm.GetState()][currentSymbol]
		)
		// print step
		fmt.Println(getTMString(tm, step) + getOutputString(action))
		// Go to Next: change state and head-position
		changeTMState(tm, action)
		step++
	}
	fmt.Println(getTMString(tm, step) + " END")

	if step >= STEP_MAX {
		return 1
	}
	return 0
}

func (svr *Server) InitializeTapeByString(tapeContent string) {
	content := []byte(tapeContent) // string 2 []byte
	svr.TuringMachine.Tape = &pb.TuringMachine_Tape{
		Cell: make([]*pb.TuringMachine_Tape_Cell, 0),
	}
	// notice: at 1st time, svr.TuringMachine.Tape == nil
	cellList := svr.TuringMachine.GetTape().GetCell()
	for coord, byteSymbol := range content {
		var cell = &pb.TuringMachine_Tape_Cell{
			Coord:  int64(coord),
			Symbol: string(byteSymbol),
		}
		cellList = append(cellList, cell)
	}
	svr.TuringMachine.GetTape().Cell = cellList
	svr.TuringMachine.HeadPosition = 1
	svr.TuringMachine.State = 0
}
