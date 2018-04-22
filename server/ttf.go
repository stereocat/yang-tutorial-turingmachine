package tm_server

import (
	pb "../proto"
    "fmt"
)

// symbol-output map
type TtfOutputMap map[string]*pb.TuringMachine_TransitionFunction_Delta_Output

// transition table function
// state-[symbol-output] map
type TTF map[uint32]TtfOutputMap

func NewTTF(ttf_config *pb.TuringMachine_TransitionFunction) TTF {
    deltaList := ttf_config.GetDelta()
	ttf := make(TTF)
	for _, delta := range deltaList {
		var input = delta.GetInput()
		if ttf[input.GetState()] == nil {
			ttf[input.GetState()] = make(TtfOutputMap)
		}
		ttf[input.GetState()][input.GetSymbol()] = delta.GetOutput()
	}
	return ttf
}

func (ttf TTF) PrintTable() {
	fmt.Printf("input        | output\n")
	fmt.Printf("state symbol | state symbol headmove\n")
	for inputState, outputMap := range ttf {
		for inputSymbol, output := range outputMap {
			fmt.Printf("   S%d %6s |    S%d %6s %8s\n",
				inputState, inputSymbol,
				output.GetState(),
				output.GetSymbol(),
				output.GetHeadMove())
		}
	}
}

func (ttf TTF) GetFinishState() uint32 {
	// in this program,
	// it assumes state of Turing Machine start by 0
	// and finished by max value of states.
	var maxState uint32 = 0
	for _, outputMap := range ttf {
		for _, output := range outputMap {
			if output.GetState() > maxState {
				maxState = output.GetState()
			}
		}
	}
	return maxState
}
