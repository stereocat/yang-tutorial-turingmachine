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

	// construct transition function table
	var tftXmlString = turingmachine.ReadTransitionTable(*transitionFunctionFileOpt)
	var tftXmlStruct = turingmachine.ParseTransitionTableString(tftXmlString)
	// turingmachine.PrintTransitionTableXmlByStruct(tftXmlStruct)
	turingmachine.CreateTransitionTable(tftXmlStruct)

	// construct rpc initial data
	var iniXmlString = turingmachine.ReadTransitionTable(*initializeFileOpt)
	var iniXmlStruct = turingmachine.ParseRpcInitializeString(iniXmlString)

	fmt.Println("#### rpc initialize string")
	fmt.Println(iniXmlString)
	fmt.Println("#### rpc parsed data")
	fmt.Println(iniXmlStruct)

	fmt.Println("## Initial")
	turingmachine.PrintRpcInitXmlByStruct(iniXmlStruct)
	turingmachine.InitializeTuringMachine(iniXmlStruct)
	turingmachine.PrintTMStateXmlByStruct(turingmachine.TMState)

	fmt.Println("## Run")
	turingmachine.RunTuringMachine()
	turingmachine.PrintTMStateXmlByStruct(turingmachine.TMState)
}
