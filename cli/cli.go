package cli

import (
	"../turingmachine"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CommandDef struct {
	Description string
	Action      func(string)
}

type CommandMap map[string]CommandDef

var CommandTable CommandMap
var scanner = bufio.NewScanner(os.Stdin)

func newCommandMap() CommandMap {
	var ct = make(CommandMap)
	ct["help"] = CommandDef{
		Description: "Print Help Message",
		Action:      printCommandHelp,
	}
	ct["?"] = ct["help"]
	ct["exit"] = CommandDef{
		Description: "Exit",
		Action:      exitCli,
	}
	ct["config"] = CommandDef{
		Description: "Read Transition Table XML",
		Action:      readTtfXml,
	}
	ct["show"] = CommandDef{
		Description: "Show Turing Machine State",
		Action:      printState,
	}
	ct["initialize"] = CommandDef{
		Description: "Read RPC Init XML",
		Action:      readRisXml,
	}
	ct["run"] = CommandDef{
		Description: "Run Turing Machine",
		Action:      run,
	}
	ct["get"] = CommandDef{
		Description: "Get Turing Machine State and Config",
		Action:      get,
	}
	return ct
}

func Start() {
	CommandTable = newCommandMap()
	var scanOk = true
	for scanOk {
		fmt.Printf("command: ")
		if scanOk = scanner.Scan(); scanOk == false {
			os.Exit(0)
		}
		var line = scanner.Text()
		if val, ok := CommandTable[line]; ok {
			val.Action(line)
		} else if line == "" {
			continue
		} else {
			printCommandError(line)
		}
	}
}

func run(_ string) {
	if turingmachine.EnableToRun() {
		var notify = turingmachine.TMState.Run()
		notify.PrintXml()
	} else {
		fmt.Println("!! Error: Transition table and/or Tape does not initialized.")
	}
}

func get(_ string) {
	turingmachine.GetConfig().PrintXml()
}

func readTtfXml(_ string) {
	turingmachine.ReadTransitionTableFromString(readXmlStringFromStdin())
}

func readRisXml(_ string) {
	turingmachine.ReadRpcInitFromString(readXmlStringFromStdin())
}

func readXmlStringFromStdin() string {
	var line string
	var lines = make([]string, 0) // multiple lines

	fmt.Println("# Paste XML data and \"EOF\"[RETURN] to finish input")
	for scanner.Scan() && line != "EOF" {
		line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n")
}

func printState(_ string) {
	fmt.Println("# Transition Table")
	if turingmachine.HasTransitionTable() {
		turingmachine.TransitionTable.PrintTable()
	} else {
		fmt.Println("## Transition Table does not initialized")
	}

	fmt.Println("# Tape")
	if turingmachine.HasTape() {
		turingmachine.TMState.Print()
	} else {
		fmt.Println("## Tape does not initialized")
	}
}

func printCommandHelp(_ string) {
	fmt.Println("    Commands:")
	for cmd, cmdDef := range CommandTable {
		fmt.Printf("\t%s\t\t%s\n", cmd, cmdDef.Description)
	}
}

func printCommandError(cmd string) {
	fmt.Printf("!! Error: unknown command: %s\n", cmd)
}

func exitCli(_ string) {
	os.Exit(0)
}
