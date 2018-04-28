package tmclient

import (
	pb "../proto"
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"regexp"
	"strings"
)

// CommandDef used for CLI command and corresponding action (function)
type CommandDef struct {
	Description string
	Action      func(*TMClient, string)
}

// CommandMap is CLI command to action string
type CommandMap map[string]CommandDef

var commandTable CommandMap
var scanner = bufio.NewScanner(os.Stdin)

// regexp of command line separator
var reSep = regexp.MustCompile(`\s+`)

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
		Action:      sendConfig,
	}
	ct["init"] = CommandDef{
		Description: "Read RPC Init XML",
		Action:      sendInitRequest,
	}
	ct["run"] = CommandDef{
		Description: "Run Turing Machine",
		Action:      run,
	}
	ct["get"] = CommandDef{
		Description: "Get Turing Machine State",
		Action:      get,
	}
	return ct
}

// TMClient is Turing Machine Client
type TMClient struct {
	TtfFileName  string
	InitFileName string
	Ctx          context.Context
	Client       pb.TuringMachineRpcClient
	UseJSON      bool
}

// NewTMClient create Turing Machine Client (Constructor)
func NewTMClient(
	ctx context.Context, client pb.TuringMachineRpcClient,
	ttfFileName string, initFileName string, useJSON bool) *TMClient {
	return &TMClient{
		TtfFileName:  ttfFileName,
		InitFileName: initFileName,
		Ctx:          ctx,
		Client:       client,
		UseJSON:      useJSON,
	}
}

// StartCli start CLI loop
func (tmClient *TMClient) StartCli() {
	commandTable = newCommandMap()
	var scanOk = true
	for scanOk {
		fmt.Printf("command> ")
		if scanOk = scanner.Scan(); scanOk == false {
			os.Exit(0)
		}
		line := scanner.Text()
		terms := reSep.Split(line, -1)
		if val, ok := commandTable[terms[0]]; ok {
			if len(terms) > 1 {
				val.Action(tmClient, terms[1])
			} else {
				val.Action(tmClient, terms[0])
			}
		} else if line == "" {
			continue
		} else {
			printCommandError(tmClient, line)
		}
	}
}

func run(tmClient *TMClient, _ string) {
	tmClient.SendRun()
}

func get(tmClient *TMClient, _ string) {
	tmClient.SendGetState()
}

func sendConfig(tmClient *TMClient, fileName string) {
	tm := tmClient.ReadTuringMachineFromFile(fileName)
	tmClient.SendConfig(tm)
}

func sendInitRequest(tmClient *TMClient, fileName string) {
	initRequest := tmClient.ReadInitRequestFromFile(fileName)
	tmClient.SendInit(initRequest)
}

func readXMLStringFromStdin() string {
	var line string
	var lines = make([]string, 0) // multiple lines

	fmt.Println("# Paste XML data and \"EOF\"[RETURN] to finish input")
	for scanner.Scan() && line != "EOF" {
		line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n")
}

func printCommandHelp(_ *TMClient, _ string) {
	fmt.Println("    Commands:")
	for cmd, cmdDef := range commandTable {
		fmt.Printf("\t%s\t\t%s\n", cmd, cmdDef.Description)
	}
}

func printCommandError(_ *TMClient, cmd string) {
	fmt.Printf("!! Error: unknown command: %s\n", cmd)
}

func exitCli(_ *TMClient, _ string) {
	os.Exit(0)
}
