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

	// read data(xml) files as []string (line-by-line)
	fmt.Println("# Read data files")
	turingmachine.ReadTransitionTableFromFile(*transitionFunctionFileOpt)
	turingmachine.ReadRpcInitFromFile(*initializeFileOpt)

	fmt.Println("# Initialize Turing Machine")
	turingmachine.PrintTMStateXml()

	fmt.Println("# Run")
	turingmachine.RunTuringMachine()
	turingmachine.PrintTMStateXml()
}
