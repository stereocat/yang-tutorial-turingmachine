package turingmachine

type OperationState struct {
	HasTtfData  bool
	HasTapeData bool
}

var opState = OperationState{
	HasTtfData:  false,
	HasTapeData: false,
}

func doneTransitionTable() {
	opState.HasTtfData = true
}

func doneTapeInitialize() {
	opState.HasTapeData = true
}

func HasTransitionTable() bool {
	return opState.HasTtfData
}

func HasTape() bool {
	return opState.HasTapeData
}

func EnableToRun() bool {
	return opState.HasTtfData && opState.HasTapeData
}
