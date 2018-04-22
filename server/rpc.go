package tm_server

import (
	pb "../proto"
	context "golang.org/x/net/context"
	"log"
)

type Server struct {
	TuringMachine   *pb.TuringMachine
	TransitionTable TTF
}

func (svr *Server) Initialize(ctx context.Context, req *pb.InitializeRequest) (*pb.Empty, error) {
	tapeContent := req.GetTapeContent()
	svr.InitializeTapeByString(tapeContent)
	log.Printf("Initialize: TapeContent: %s\n", tapeContent)
	return &pb.Empty{}, nil
}

func (svr *Server) Configure(ctx context.Context, req *pb.Config) (*pb.Empty, error) {
	svr.TuringMachine = req.GetTuringMachine()
	log.Printf("Config: TuringMachine: %s\n", svr.TuringMachine.String())
	svr.TransitionTable = NewTTF(svr.TuringMachine.GetTransitionFunction())
	svr.TransitionTable.PrintTable()
	return &pb.Empty{}, nil
}

func (svr *Server) Run(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	svr.RunTM()
	return &pb.Empty{}, nil
}
