package tmserver

import (
	pb "../proto"
	"fmt"
)

// TtfDeltaMap is part of transition table function.
// symbol-output map
type TtfDeltaMap map[string]*pb.TuringMachine_TransitionFunction_Delta

// TTF is transition table function
// state-[symbol-output] map
type TTF map[uint32]TtfDeltaMap

// NewTTF is Constructor
func NewTTF(ttfConfig *pb.TuringMachine_TransitionFunction) TTF {
	ttf := make(TTF)
	for _, delta := range ttfConfig.GetDelta() {
		input := delta.GetInput()
		if ttf[input.GetState()] == nil {
			// create space of pointer
			ttf[input.GetState()] = make(TtfDeltaMap)
		}
		ttf[input.GetState()][input.GetSymbol()] = delta
	}
	return ttf
}

// Print Transition Table to Stdout
func (ttf TTF) Print() {
	fmt.Printf("input        | output\n")
	fmt.Printf("state symbol | state symbol move\n")
	for inputState, deltaMap := range ttf {
		for inputSymbol, delta := range deltaMap {
			fmt.Printf("   S%d %6s |  %s %s\n",
				inputState, inputSymbol,
				delta.GetOutput().ToString(), delta.GetLabel())
		}
	}
}

// GetFinishState return Finish (maximum) state from Transition Table
func (ttf TTF) GetFinishState() uint32 {
	// in this program,
	// it assumes state of Turing Machine start by 0
	// and finished by max value of states.
	var maxState uint32 // 0 (default)
	for _, deltaMap := range ttf {
		for _, delta := range deltaMap {
			if delta.GetOutput().GetState() > maxState {
				maxState = delta.GetOutput().GetState()
			}
		}
	}
	return maxState
}
