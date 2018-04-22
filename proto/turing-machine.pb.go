// Code generated by protoc-gen-go. DO NOT EDIT.
// source: turing-machine.proto

/*
Package turing_machine is a generated protocol buffer package.

Data model for the Turing Machine.

It is generated from these files:
	turing-machine.proto

It has these top-level messages:
	InitializeRequest
	Empty
	Halted
	Run
	Config
	TuringMachine
*/
package turing_machine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Move the head one cell to the left or right
type HeadMove int32

const (
	HeadMove_LEFT  HeadMove = 0
	HeadMove_RIGHT HeadMove = 1
)

var HeadMove_name = map[int32]string{
	0: "LEFT",
	1: "RIGHT",
}
var HeadMove_value = map[string]int32{
	"LEFT":  0,
	"RIGHT": 1,
}

func (x HeadMove) String() string {
	return proto.EnumName(HeadMove_name, int32(x))
}
func (HeadMove) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InitializeRequest struct {
	// The string with which the tape shall be initialized. The
	// leftmost symbol will be at tape coordinate 0.
	TapeContent string `protobuf:"bytes,1,opt,name=tape_content,json=tapeContent" json:"tape_content,omitempty"`
}

func (m *InitializeRequest) Reset()                    { *m = InitializeRequest{} }
func (m *InitializeRequest) String() string            { return proto.CompactTextString(m) }
func (*InitializeRequest) ProtoMessage()               {}
func (*InitializeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InitializeRequest) GetTapeContent() string {
	if m != nil {
		return m.TapeContent
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The Turing Machine has halted. This means that there is no
// transition rule for the current state and tape symbol.
type Halted struct {
	// The state of the control unit in which the machine has
	// halted.
	State uint32 `protobuf:"varint,1,opt,name=state" json:"state,omitempty"`
}

func (m *Halted) Reset()                    { *m = Halted{} }
func (m *Halted) String() string            { return proto.CompactTextString(m) }
func (*Halted) ProtoMessage()               {}
func (*Halted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Halted) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

// Start the Turing Machine operation.
type Run struct {
}

func (m *Run) Reset()                    { *m = Run{} }
func (m *Run) String() string            { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()               {}
func (*Run) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Config data
type Config struct {
	// @inject_tag: xml:"xmlns,attr"
	Xmlns string `protobuf:"bytes,1,opt,name=xmlns" json:"xmlns,omitempty" xml:"xmlns,attr"`
	// @inject_tag: xml:"turing-machine"
	TuringMachine *TuringMachine `protobuf:"bytes,2,opt,name=turing_machine,json=turingMachine" json:"turing_machine,omitempty" xml:"turing-machine"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Config) GetXmlns() string {
	if m != nil {
		return m.Xmlns
	}
	return ""
}

func (m *Config) GetTuringMachine() *TuringMachine {
	if m != nil {
		return m.TuringMachine
	}
	return nil
}

// State data and configuration of a Turing Machine.
type TuringMachine struct {
	// Position of tape read/write head.
	// @inject_tag: xml:"head-position"
	HeadPosition int64 `protobuf:"varint,1,opt,name=head_position,json=headPosition" json:"head_position,omitempty" xml:"head-position"`
	// Current state of the control unit.
	// The initial state is 0.
	// @inject_tag: xml:"state"
	State uint32 `protobuf:"varint,2,opt,name=state" json:"state,omitempty" xml:"state"`
	// @inject_tag: xml:"tape"
	Tape *TuringMachine_Tape `protobuf:"bytes,3,opt,name=tape" json:"tape,omitempty" xml:"tape"`
	// @inject_tag: xml:"transition-function"
	TransitionFunction *TuringMachine_TransitionFunction `protobuf:"bytes,4,opt,name=transition_function,json=transitionFunction" json:"transition_function,omitempty" xml:"transition-function"`
}

func (m *TuringMachine) Reset()                    { *m = TuringMachine{} }
func (m *TuringMachine) String() string            { return proto.CompactTextString(m) }
func (*TuringMachine) ProtoMessage()               {}
func (*TuringMachine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TuringMachine) GetHeadPosition() int64 {
	if m != nil {
		return m.HeadPosition
	}
	return 0
}

func (m *TuringMachine) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *TuringMachine) GetTape() *TuringMachine_Tape {
	if m != nil {
		return m.Tape
	}
	return nil
}

func (m *TuringMachine) GetTransitionFunction() *TuringMachine_TransitionFunction {
	if m != nil {
		return m.TransitionFunction
	}
	return nil
}

// The contents of the tape.
type TuringMachine_Tape struct {
	// @inject_tag: xml:"cell"
	Cell []*TuringMachine_Tape_Cell `protobuf:"bytes,1,rep,name=cell" json:"cell,omitempty" xml:"cell"`
}

func (m *TuringMachine_Tape) Reset()                    { *m = TuringMachine_Tape{} }
func (m *TuringMachine_Tape) String() string            { return proto.CompactTextString(m) }
func (*TuringMachine_Tape) ProtoMessage()               {}
func (*TuringMachine_Tape) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *TuringMachine_Tape) GetCell() []*TuringMachine_Tape_Cell {
	if m != nil {
		return m.Cell
	}
	return nil
}

// List of non-blank cells.
type TuringMachine_Tape_Cell struct {
	// Coordinate (index) of the tape cell.
	// @inject_tag: xml:"coord"
	Coord int64 `protobuf:"varint,1,opt,name=coord" json:"coord,omitempty" xml:"coord"`
	// Symbol appearing in the tape cell.
	// Blank (empty string) is not allowed here because the
	// 'cell' list only contains non-blank cells.
	// @inject_tag: xml:"symbol"
	Symbol string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty" xml:"symbol"`
}

func (m *TuringMachine_Tape_Cell) Reset()                    { *m = TuringMachine_Tape_Cell{} }
func (m *TuringMachine_Tape_Cell) String() string            { return proto.CompactTextString(m) }
func (*TuringMachine_Tape_Cell) ProtoMessage()               {}
func (*TuringMachine_Tape_Cell) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0, 0} }

func (m *TuringMachine_Tape_Cell) GetCoord() int64 {
	if m != nil {
		return m.Coord
	}
	return 0
}

func (m *TuringMachine_Tape_Cell) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

// The Turing Machine is configured by specifying the
// transition function.
type TuringMachine_TransitionFunction struct {
	// @inject_tag: xml:"delta"
	Delta []*TuringMachine_TransitionFunction_Delta `protobuf:"bytes,1,rep,name=delta" json:"delta,omitempty" xml:"delta"`
}

func (m *TuringMachine_TransitionFunction) Reset()         { *m = TuringMachine_TransitionFunction{} }
func (m *TuringMachine_TransitionFunction) String() string { return proto.CompactTextString(m) }
func (*TuringMachine_TransitionFunction) ProtoMessage()    {}
func (*TuringMachine_TransitionFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1}
}

func (m *TuringMachine_TransitionFunction) GetDelta() []*TuringMachine_TransitionFunction_Delta {
	if m != nil {
		return m.Delta
	}
	return nil
}

// The list of transition rules.
type TuringMachine_TransitionFunction_Delta struct {
	// @inject_tag: xml:"input"
	Input *TuringMachine_TransitionFunction_Delta_Input `protobuf:"bytes,1,opt,name=input" json:"input,omitempty" xml:"input"`
	// An arbitrary label of the transition rule.
	// @inject_tag: xml:"label"
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty" xml:"label"`
	// @inject_tag: xml:"output"
	Output *TuringMachine_TransitionFunction_Delta_Output `protobuf:"bytes,3,opt,name=output" json:"output,omitempty" xml:"output"`
}

func (m *TuringMachine_TransitionFunction_Delta) Reset() {
	*m = TuringMachine_TransitionFunction_Delta{}
}
func (m *TuringMachine_TransitionFunction_Delta) String() string { return proto.CompactTextString(m) }
func (*TuringMachine_TransitionFunction_Delta) ProtoMessage()    {}
func (*TuringMachine_TransitionFunction_Delta) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1, 0}
}

func (m *TuringMachine_TransitionFunction_Delta) GetInput() *TuringMachine_TransitionFunction_Delta_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *TuringMachine_TransitionFunction_Delta) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *TuringMachine_TransitionFunction_Delta) GetOutput() *TuringMachine_TransitionFunction_Delta_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

// Input parameters (arguments) of the transition rule.
type TuringMachine_TransitionFunction_Delta_Input struct {
	// Current state of the control unit.
	// @inject_tag: xml:"state"
	State uint32 `protobuf:"varint,1,opt,name=state" json:"state,omitempty" xml:"state"`
	// Symbol read from the tape cell.
	// @inject_tag: xml:"symbol"
	Symbol string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty" xml:"symbol"`
}

func (m *TuringMachine_TransitionFunction_Delta_Input) Reset() {
	*m = TuringMachine_TransitionFunction_Delta_Input{}
}
func (m *TuringMachine_TransitionFunction_Delta_Input) String() string {
	return proto.CompactTextString(m)
}
func (*TuringMachine_TransitionFunction_Delta_Input) ProtoMessage() {}
func (*TuringMachine_TransitionFunction_Delta_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1, 0, 0}
}

func (m *TuringMachine_TransitionFunction_Delta_Input) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *TuringMachine_TransitionFunction_Delta_Input) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

// Output values of the transition rule.
type TuringMachine_TransitionFunction_Delta_Output struct {
	// HeadMove head_move = 1;
	// @inject_tag: xml:"head-move"
	HeadMove string `protobuf:"bytes,1,opt,name=head_move,json=headMove" json:"head_move,omitempty" xml:"head-move"`
	// New state of the control unit. If this leaf is not
	// present, the state doesn't change.
	// @inject_tag: xml:"state"
	State uint32 `protobuf:"varint,2,opt,name=state" json:"state,omitempty" xml:"state"`
	// Symbol to be written to the tape cell. If this leaf is
	// not present, the symbol doesn't change.
	// @inject_tag: xml:"symbol"
	Symbol string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty" xml:"symbol"`
}

func (m *TuringMachine_TransitionFunction_Delta_Output) Reset() {
	*m = TuringMachine_TransitionFunction_Delta_Output{}
}
func (m *TuringMachine_TransitionFunction_Delta_Output) String() string {
	return proto.CompactTextString(m)
}
func (*TuringMachine_TransitionFunction_Delta_Output) ProtoMessage() {}
func (*TuringMachine_TransitionFunction_Delta_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1, 0, 1}
}

func (m *TuringMachine_TransitionFunction_Delta_Output) GetHeadMove() string {
	if m != nil {
		return m.HeadMove
	}
	return ""
}

func (m *TuringMachine_TransitionFunction_Delta_Output) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *TuringMachine_TransitionFunction_Delta_Output) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func init() {
	proto.RegisterType((*InitializeRequest)(nil), "turing_machine.InitializeRequest")
	proto.RegisterType((*Empty)(nil), "turing_machine.Empty")
	proto.RegisterType((*Halted)(nil), "turing_machine.Halted")
	proto.RegisterType((*Run)(nil), "turing_machine.Run")
	proto.RegisterType((*Config)(nil), "turing_machine.Config")
	proto.RegisterType((*TuringMachine)(nil), "turing_machine.TuringMachine")
	proto.RegisterType((*TuringMachine_Tape)(nil), "turing_machine.TuringMachine.Tape")
	proto.RegisterType((*TuringMachine_Tape_Cell)(nil), "turing_machine.TuringMachine.Tape.Cell")
	proto.RegisterType((*TuringMachine_TransitionFunction)(nil), "turing_machine.TuringMachine.TransitionFunction")
	proto.RegisterType((*TuringMachine_TransitionFunction_Delta)(nil), "turing_machine.TuringMachine.TransitionFunction.Delta")
	proto.RegisterType((*TuringMachine_TransitionFunction_Delta_Input)(nil), "turing_machine.TuringMachine.TransitionFunction.Delta.Input")
	proto.RegisterType((*TuringMachine_TransitionFunction_Delta_Output)(nil), "turing_machine.TuringMachine.TransitionFunction.Delta.Output")
	proto.RegisterEnum("turing_machine.HeadMove", HeadMove_name, HeadMove_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TuringMachineRpc service

type TuringMachineRpcClient interface {
	Configure(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error)
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*Empty, error)
	Run(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Halted, error)
}

type turingMachineRpcClient struct {
	cc *grpc.ClientConn
}

func NewTuringMachineRpcClient(cc *grpc.ClientConn) TuringMachineRpcClient {
	return &turingMachineRpcClient{cc}
}

func (c *turingMachineRpcClient) Configure(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/turing_machine.TuringMachineRpc/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turingMachineRpcClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/turing_machine.TuringMachineRpc/Initialize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turingMachineRpcClient) Run(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Halted, error) {
	out := new(Halted)
	err := grpc.Invoke(ctx, "/turing_machine.TuringMachineRpc/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TuringMachineRpc service

type TuringMachineRpcServer interface {
	Configure(context.Context, *Config) (*Empty, error)
	Initialize(context.Context, *InitializeRequest) (*Empty, error)
	Run(context.Context, *Empty) (*Halted, error)
}

func RegisterTuringMachineRpcServer(s *grpc.Server, srv TuringMachineRpcServer) {
	s.RegisterService(&_TuringMachineRpc_serviceDesc, srv)
}

func _TuringMachineRpc_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TuringMachineRpcServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turing_machine.TuringMachineRpc/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TuringMachineRpcServer).Configure(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _TuringMachineRpc_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TuringMachineRpcServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turing_machine.TuringMachineRpc/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TuringMachineRpcServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TuringMachineRpc_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TuringMachineRpcServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turing_machine.TuringMachineRpc/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TuringMachineRpcServer).Run(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TuringMachineRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "turing_machine.TuringMachineRpc",
	HandlerType: (*TuringMachineRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _TuringMachineRpc_Configure_Handler,
		},
		{
			MethodName: "Initialize",
			Handler:    _TuringMachineRpc_Initialize_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _TuringMachineRpc_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "turing-machine.proto",
}

func init() { proto.RegisterFile("turing-machine.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xf1, 0x85, 0x64, 0xd2, 0xa0, 0x30, 0x94, 0x2a, 0x32, 0x02, 0x52, 0xf3, 0x40, 0x84,
	0x84, 0x85, 0x42, 0xc9, 0x03, 0x97, 0xa7, 0xb4, 0x21, 0x91, 0x5a, 0x81, 0x96, 0xf0, 0x1c, 0x6d,
	0x9c, 0x6d, 0x6b, 0xc9, 0xd9, 0x35, 0xc9, 0xba, 0x22, 0x7c, 0x0b, 0x3f, 0xc2, 0x3f, 0xf0, 0xc8,
	0x07, 0xa1, 0xdd, 0x75, 0xc8, 0xc5, 0x84, 0x8a, 0x3e, 0x9e, 0xe3, 0x99, 0x33, 0x67, 0xc6, 0x3b,
	0x03, 0xfb, 0x32, 0x9b, 0xc5, 0xfc, 0xe2, 0xf9, 0x94, 0x46, 0x97, 0x31, 0x67, 0x61, 0x3a, 0x13,
	0x52, 0xe0, 0x1d, 0xc3, 0x8e, 0x72, 0x36, 0xe8, 0xc0, 0xdd, 0x01, 0x8f, 0x65, 0x4c, 0x93, 0xf8,
	0x1b, 0x23, 0xec, 0x4b, 0xc6, 0xe6, 0x12, 0x0f, 0x61, 0x4f, 0xd2, 0x94, 0x8d, 0x22, 0xc1, 0x25,
	0xe3, 0xb2, 0x61, 0x35, 0xad, 0x56, 0x85, 0x54, 0x15, 0xd7, 0x35, 0x54, 0x70, 0x1b, 0xdc, 0x93,
	0x69, 0x2a, 0x17, 0xc1, 0x23, 0xf0, 0xfa, 0x34, 0x91, 0x6c, 0x82, 0xfb, 0xe0, 0xce, 0x25, 0x95,
	0x4c, 0x87, 0xd7, 0x88, 0x01, 0x81, 0x0b, 0x36, 0xc9, 0x78, 0x30, 0x01, 0xaf, 0x2b, 0xf8, 0x79,
	0x7c, 0xa1, 0xc2, 0xbe, 0x4e, 0x13, 0x3e, 0xcf, 0x55, 0x0d, 0xc0, 0x63, 0xd8, 0x72, 0xd6, 0x28,
	0x35, 0xad, 0x56, 0xb5, 0xfd, 0x30, 0xdc, 0xa4, 0xc3, 0xa1, 0x86, 0x67, 0x06, 0x91, 0x9a, 0x5c,
	0x87, 0xc1, 0x77, 0x0f, 0x6a, 0x1b, 0x01, 0xf8, 0x04, 0x6a, 0x97, 0x8c, 0x4e, 0x46, 0xa9, 0x98,
	0xc7, 0x32, 0x16, 0x5c, 0x57, 0xb5, 0xc9, 0x9e, 0x22, 0x3f, 0xe6, 0xdc, 0xca, 0x79, 0x69, 0xcd,
	0x39, 0x76, 0xc0, 0x51, 0x1d, 0x37, 0x6c, 0x6d, 0x24, 0xf8, 0xa7, 0x91, 0x70, 0x48, 0x53, 0x46,
	0x74, 0x3c, 0x52, 0xb8, 0x27, 0x67, 0x94, 0x1b, 0xed, 0xd1, 0x79, 0xc6, 0x23, 0x5d, 0xd8, 0xd1,
	0x32, 0x2f, 0xae, 0x91, 0xf9, 0x93, 0xd8, 0xcb, 0xf3, 0x08, 0xca, 0x02, 0xe7, 0x2f, 0xc0, 0x51,
	0x05, 0xf1, 0x0d, 0x38, 0x11, 0x4b, 0x92, 0x86, 0xd5, 0xb4, 0x5b, 0xd5, 0xf6, 0xd3, 0xeb, 0x2d,
	0x86, 0x5d, 0x96, 0x24, 0x44, 0x27, 0xf9, 0x47, 0xe0, 0x28, 0xa4, 0xba, 0x8f, 0x84, 0x98, 0x4d,
	0xf2, 0xd1, 0x18, 0x80, 0x07, 0xe0, 0xcd, 0x17, 0xd3, 0xb1, 0x48, 0xf4, 0x50, 0x2a, 0x24, 0x47,
	0xfe, 0x0f, 0x1b, 0xb0, 0xe8, 0x12, 0x4f, 0xc1, 0x9d, 0xb0, 0x44, 0xd2, 0xdc, 0x4a, 0xe7, 0x7f,
	0xdb, 0x0c, 0x8f, 0x55, 0x36, 0x31, 0x22, 0xfe, 0xaf, 0x12, 0xb8, 0x9a, 0x40, 0x02, 0x6e, 0xcc,
	0xd3, 0xcc, 0xbc, 0xc1, 0x6a, 0xfb, 0xed, 0xcd, 0x74, 0xc3, 0x81, 0xd2, 0x20, 0x46, 0x4a, 0x35,
	0x9c, 0xd0, 0x31, 0x5b, 0x76, 0x66, 0x00, 0x7e, 0x06, 0x4f, 0x64, 0x52, 0x95, 0x32, 0x3f, 0xfc,
	0xdd, 0x0d, 0x4b, 0x7d, 0xd0, 0x22, 0x24, 0x17, 0xf3, 0x5f, 0x81, 0x3b, 0x58, 0x56, 0x2d, 0xae,
	0xc7, 0xce, 0x31, 0x7f, 0x02, 0xcf, 0x08, 0xe1, 0x03, 0xa8, 0xe8, 0x17, 0x3c, 0x15, 0x57, 0x2c,
	0xdf, 0x99, 0xb2, 0x22, 0xce, 0xc4, 0x15, 0xdb, 0xf1, 0x72, 0x57, 0xa2, 0xf6, 0xba, 0xe8, 0xb3,
	0xc7, 0x50, 0xee, 0x2f, 0x33, 0xcb, 0xe0, 0x9c, 0x9e, 0xf4, 0x86, 0xf5, 0x5b, 0x58, 0x01, 0x97,
	0x0c, 0xde, 0xf7, 0x87, 0x75, 0xab, 0xfd, 0xd3, 0x82, 0xfa, 0xe6, 0x82, 0xa5, 0x11, 0xbe, 0x86,
	0x8a, 0x59, 0xdd, 0x6c, 0xc6, 0xf0, 0x60, 0x7b, 0x2a, 0xe6, 0x93, 0x7f, 0x7f, 0x9b, 0xd7, 0xd7,
	0x01, 0x7b, 0x00, 0xab, 0xf3, 0x82, 0x87, 0xdb, 0x41, 0x85, 0xd3, 0xb3, 0x4b, 0xe7, 0x48, 0x5f,
	0x11, 0xfc, 0xfb, 0x57, 0xbf, 0x60, 0xca, 0x5c, 0xa4, 0xb1, 0xa7, 0x6f, 0xde, 0xcb, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf7, 0xa5, 0x83, 0x50, 0x0b, 0x05, 0x00, 0x00,
}
