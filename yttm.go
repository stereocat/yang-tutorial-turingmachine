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
	var tftXmlString = turingmachine.ReadDataFile(*transitionFunctionFileOpt)
	var iniXmlString = turingmachine.ReadDataFile(*initializeFileOpt)

	// construct transition function table
	fmt.Println("# parse transition function data")
	var tftXmlStruct = turingmachine.ParseTransitionTableString(tftXmlString)
	// fmt.Println("## xml data")
	// turingmachine.PrintTransitionTableXmlByStruct(tftXmlStruct)
	fmt.Println("# create transition function")
	turingmachine.CreateTransitionTable(tftXmlStruct)

	// construct rpc initial data
	fmt.Println("# parse rpc initial data")
	var iniXmlStruct = turingmachine.ParseRpcInitializeString(iniXmlString)

	fmt.Println("## rpc initialize string")
	fmt.Println(iniXmlString)
	fmt.Println("## rpc parsed data")
	fmt.Println(iniXmlStruct)

	fmt.Println("# Initialize Turing Machine")
	turingmachine.PrintRpcInitXmlByStruct(iniXmlStruct)
	turingmachine.InitializeTuringMachine(iniXmlStruct)
	turingmachine.PrintTMStateXmlByStruct(turingmachine.TMState)

	fmt.Println("# Run")
	turingmachine.RunTuringMachine()
	turingmachine.PrintTMStateXmlByStruct(turingmachine.TMState)
}
