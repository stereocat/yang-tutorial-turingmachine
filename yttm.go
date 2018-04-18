package main

import (
	"./cli"
	"./turingmachine"
	"flag"
	"fmt"
)

// Options
var (
	interactive               = flag.Bool("c", false, "CLI (interactive)")
	verbose                   = flag.Bool("v", false, "Verbose output")
	transitionFunctionFileOpt = flag.String("t", "", "transition function table xml")
	initializeFileOpt         = flag.String("i", "", "rpc initialize xml")
)

func main() {
	flag.Parse()

	if *verbose {
		turingmachine.SetVerbose(true)
	}

	if *transitionFunctionFileOpt != "" {
		turingmachine.ReadTransitionTableFromFile(*transitionFunctionFileOpt)
	}
	if *initializeFileOpt != "" {
		turingmachine.ReadRpcInitFromFile(*initializeFileOpt)
	}

	if *interactive {
		cli.Start()
	} else {
		fmt.Println("# Initialize Turing Machine")
		turingmachine.TMState.PrintXml()
		fmt.Println("# Run")
		turingmachine.TMState.Run()
		turingmachine.GetConfig().PrintXml()
	}
}
