package tmserver

import (
	pb "../proto"
	"golang.org/x/net/context"
	"log"
)

// TMServer has state of Turing Machine
//   - TuringMachine: data received from RPC
//   - TransitionTable: struct build from TuringMachine
type TMServer struct {
	TuringMachine   *pb.TuringMachine
	TransitionTable TTF
}

// Initialize is gRPC Interface to receive tape content
// This method initialize Turing Machine head-position/state
func (svr *TMServer) Initialize(ctx context.Context, req *pb.InitializeRequest) (*pb.Empty, error) {
	tapeContent := req.GetTapeContent()
	svr.InitializeTapeByString(tapeContent)
	log.Printf("Initialize: TapeContent: %s\n", tapeContent)
	return &pb.Empty{}, nil
}

// Configure is gRPC Interface to receive transion table function
func (svr *TMServer) Configure(ctx context.Context, req *pb.Config) (*pb.Empty, error) {
	reqTm := req.GetTuringMachine()
	if reqTm != nil {
		// overwrite if found in request
		reqTtf := reqTm.GetTransitionFunction()
		if reqTtf != nil {
			svr.TuringMachine.TransitionFunction = reqTtf
		}
		reqTape := reqTm.GetTape()
		if reqTape != nil {
			svr.TuringMachine.Tape = reqTape
		}
	}
	svr.TransitionTable = NewTTF(svr.TuringMachine.GetTransitionFunction())
	svr.TransitionTable.Print()
	return &pb.Empty{}, nil
}

// Run is gRPC Interface to exec Turing Machine calculation
func (svr *TMServer) Run(ctx context.Context, _ *pb.Empty) (*pb.Halted, error) {
	var halted = &pb.Halted{State: svr.RunTM()}
	return halted, nil
}

// GetState is gRPC Interface
// to send Turing Machine state, tape content and transition table
func (svr *TMServer) GetState(ctx context.Context, _ *pb.Empty) (*pb.TuringMachine, error) {
	return svr.TuringMachine, nil
}
