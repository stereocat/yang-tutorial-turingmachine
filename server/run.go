package tmserver

import (
	pb "../proto"
	"fmt"
)

// StepMax means maximum step count to avoid infinite loop
const StepMax = 20

// RunTM execute Turing Machine Calculation
// returns 0 (normal) >1 (error)
func (svr *TMServer) RunTM() (lastState uint32, err bool) {
	var step uint32 = 1
	tm := svr.TuringMachine

	// check state
	if len(tm.GetTape().GetCell()) < 1 {
		return 0, true // Cannot run when tape is not initialized
	}
	if len(svr.TuringMachine.GetTransitionFunction().GetDelta()) < 1 {
		return 0, true // Cannot run when transition function table is not initialized
	}

	// print header
	indent := len([]byte(tm.ToString(step))) - 20 // offset
	headerString := fmt.Sprintf("Step State | Tape %%%ds | Next Write Move\n", indent)
	fmt.Printf(headerString, " ")

	// Run
	finishState := svr.TransitionTable.GetFinishState()
	for tm.GetState() != finishState && step < StepMax {
		// read symbol under head-position
		cellList := tm.GetTape().GetCell()

		// find next output by current state and symbol
		var delta *pb.TuringMachine_TransitionFunction_Delta
		headPos := tm.GetHeadPosition()
		if headPos >= 0 && int(headPos) < len(cellList) {
			delta = svr.TransitionTable[tm.GetState()][cellList[headPos].GetSymbol()]
		} else {
			// out of range
			delta = svr.TransitionTable[tm.GetState()][""]
			fmt.Println("             cell-list out-of-lange :",
				delta.GetOutput().ToString())
		}

		// print step
		fmt.Println(tm.ToString(step), delta.GetOutput().ToString(), delta.GetLabel())
		// Go to Next: change state and head-position
		tm.ChangeState(delta.GetOutput())
		step++
	}
	fmt.Println(tm.ToString(step) + " END")

	if step >= StepMax {
		return tm.GetState(), true
	}
	return tm.GetState(), false
}

// InitializeTapeByString initialize
// tape content, state, head-position of Turing Machine
func (svr *TMServer) InitializeTapeByString(tapeContent string) {
	// initialize tape by 0 length cell list
	svr.TuringMachine.Tape = &pb.TuringMachine_Tape{
		Cell: make([]*pb.TuringMachine_Tape_Cell, 0),
	}

	// notice: at 1st time, svr.TuringMachine.Tape == nil
	content := []byte(tapeContent) // convert string to []byte
	cellList := svr.TuringMachine.GetTape().GetCell()
	for coord, byteSymbol := range content {
		var cell = &pb.TuringMachine_Tape_Cell{
			Coord:  int64(coord),
			Symbol: string(byteSymbol),
		}
		cellList = append(cellList, cell)
	}
	// DO NOT initialize TransitionFunction
	svr.TuringMachine.GetTape().Cell = cellList
	svr.TuringMachine.HeadPosition = 0
	svr.TuringMachine.State = 0
}

// NewTMServer returns empty Turing Machine Server
func NewTMServer() *TMServer {
	return &TMServer{
		TuringMachine: &pb.TuringMachine{
			HeadPosition: 0,
			State:        0,
			Tape: &pb.TuringMachine_Tape{
				Cell: make([]*pb.TuringMachine_Tape_Cell, 0),
			},
			// set to avoid empty(nil) function calling
			TransitionFunction: &pb.TuringMachine_TransitionFunction{},
		},
		TransitionTable: TTF{}, // map(ref-type)
	}
}
