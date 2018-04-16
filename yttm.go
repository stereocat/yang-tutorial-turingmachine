package main

import (
	"./turingmachine"
	"flag"
	"fmt"
)

// Options
var (
	transitionFunctionFileOpt = flag.String("t", "", "transition function table xml")
	initializeFileOpt         = flag.String("i", "", "rpc initialize xml")
)

func main() {
	flag.Parse()

	fmt.Println("# Read data files")
	turingmachine.ReadTransitionTableFromFile(*transitionFunctionFileOpt)
	turingmachine.ReadRpcInitFromFile(*initializeFileOpt)

	fmt.Println("# Initialize Turing Machine")
	turingmachine.TMState.PrintXml()

	fmt.Println("# Run")
	turingmachine.TMState.Run()
	turingmachine.TMState.PrintXml()
}
