package tm_server

import (
	pb "../proto"
	"fmt"
)

const STEP_MAX = 100

func (svr *Server) RunTM() uint32 {
	tm := svr.TuringMachine
	step := 1

	// check state
	if len(tm.GetTape().GetCell()) < 1 {
		return 1 // Cannot run when tape is not initialized
	}
	if len(svr.TuringMachine.GetTransitionFunction().GetDelta()) < 1 {
		return 2 // Cannot run when transition function table is not initialized
	}

	// print header
	indent := len([]byte(tm.ToString(step))) - 20 // offset
	headerString := fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	finishState := svr.TransitionTable.GetFinishState()
	for tm.GetState() != finishState && step < STEP_MAX {
		// read symbol under head-position
		cellList := tm.GetTape().GetCell()
		currentSymbol := cellList[tm.GetHeadPosition()].GetSymbol()
		// find next action by current state and symbol
		action := svr.TransitionTable[tm.GetState()][currentSymbol]

		// print step
		fmt.Println(tm.ToString(step) + action.ToString())
		// Go to Next: change state and head-position
		tm.ChangeState(action)
		step++
	}
	fmt.Println(tm.ToString(step) + " END")

	if step >= STEP_MAX {
		return 9
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
