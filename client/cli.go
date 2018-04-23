package tm_client

import (
	pb "../proto"
	"bufio"
	"fmt"
	context "golang.org/x/net/context"
	"os"
	"strings"
)

type CommandDef struct {
	Description string
	Action      func(*ClientCli, string)
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
	// ct["show"] = CommandDef{
	// 	Description: "Show Turing Machine State",
	// 	Action:      printState,
	// }
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

type ClientCli struct {
	TtfFileName  string
	InitFileName string
	Ctx          context.Context
	Client       pb.TuringMachineRpcClient
}

func NewClientCli(ctx context.Context, client pb.TuringMachineRpcClient, ttffn string, initfn string) *ClientCli {
	return &ClientCli{
		TtfFileName:  ttffn,
		InitFileName: initfn,
		Ctx:          ctx,
		Client:       client,
	}
}

func (ccli *ClientCli) Start() {
	CommandTable = newCommandMap()
	var scanOk = true
	for scanOk {
		fmt.Printf("command: ")
		if scanOk = scanner.Scan(); scanOk == false {
			os.Exit(0)
		}
		var line = scanner.Text()
		if val, ok := CommandTable[line]; ok {
			val.Action(ccli, line)
		} else if line == "" {
			continue
		} else {
			printCommandError(ccli, line)
		}
	}
}

func run(ccli *ClientCli, _ string) {
	SendRun(ccli.Ctx, ccli.Client)
}

func get(ccli *ClientCli, _ string) {
	SendGetState(ccli.Ctx, ccli.Client)
}

func readTtfXml(ccli *ClientCli, _ string) {
	SendConfig(ccli.Ctx, ccli.Client, ccli.TtfFileName)
}

func readRisXml(ccli *ClientCli, _ string) {
	SendInit(ccli.Ctx, ccli.Client, ccli.InitFileName)
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

func printCommandHelp(_ *ClientCli, _ string) {
	fmt.Println("    Commands:")
	for cmd, cmdDef := range CommandTable {
		fmt.Printf("\t%s\t\t%s\n", cmd, cmdDef.Description)
	}
}

func printCommandError(_ *ClientCli, cmd string) {
	fmt.Printf("!! Error: unknown command: %s\n", cmd)
}

func exitCli(_ *ClientCli, _ string) {
	os.Exit(0)
}
