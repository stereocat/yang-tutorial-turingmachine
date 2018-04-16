package turingmachine

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	XMLName       xml.Name            `xml:"config"`
	TuringMachine TuringMachineConfig `xml:"turing-machine"`
}
type TuringMachineConfig struct {
	// XMLName xml.Name `xml:"turing-machine"`
	TransitionFunction TransitionFunction `xml:"transition-function"`
}
type TransitionFunction struct {
	// XMLName xml.Name `xml:"transition-function"`
	DeltaList []Delta `xml:"delta"`
}
type Delta struct {
	// XMLName xml.Name `xml:"delta"`
	Label  string `xml:"label"`
	Input  Input  `xml:"input"`
	Output Output `xml:"output"`
}
type Input struct {
	// XMLName xml.Name `xml:"input"`
	State  uint16 `xml:"state"`
	Symbol string `xml:"symbol"`
}
type Output struct {
	// XMLName xml.Name `xml:"output"`
	State    uint16 `xml:"state"`
	Symbol   string `xml:"symbol"`
	HeadMove string `xml:"head-move"`
}

var transitionTableString string
var transitionTableStruct Config

type TTF map[uint16]map[string]Output // transition table function
var TransitionTable = make(TTF)

func ReadTransitionTableFromFile(xmlFileName string) {
	// construct transition table
	transitionTableString = readXmlString(xmlFileName)
	(&transitionTableStruct).new() // pointer to rewrite self
	transitionTableStruct.printXml()
	TransitionTable.new()
}

func readXmlString(xmlFileName string) string {
	fmt.Println("#### File: ", xmlFileName)
	xmlFile, err := os.Open(xmlFileName)
	if err != nil {
		fmt.Printf("!! Error: Cannot read file:%s\n", xmlFileName)
		os.Exit(1)
	} else {
		defer xmlFile.Close() // close finished readXmlFile
	}
	return readXmlFile(xmlFile)
}

func readXmlFile(xmlFile *os.File) string {
	// read data from file
	var scanner = bufio.NewScanner(xmlFile)
	var lines = make([]string, 0) // multiple lines
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n") // convert to single line
}

func (ttsPtr *Config) new() { // pointer to rewrite self
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(transitionTableString), ttsPtr); err != nil {
		fmt.Println("!! Error: TransitionTable XML Unmarshal error: ", err)
	}
}

func (tts Config) printXml() {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tts, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}

func (transitionTable TTF) new() {
	var deltaList = transitionTableStruct.TuringMachine.TransitionFunction.DeltaList

	for _, delta := range deltaList {
		var input = delta.Input
		var output = delta.Output
		if transitionTable[input.State] == nil {
			transitionTable[input.State] = make(map[string]Output)
		}
		transitionTable[input.State][input.Symbol] = output
	}

	fmt.Printf("input        | output\n")
	fmt.Printf("state symbol | state symbol headmove\n")
	for inputState, outputMap := range transitionTable {
		for inputSymbol, output := range outputMap {
			fmt.Printf("   S%d %6s |    S%d %6s %8s\n", inputState, inputSymbol, output.State, output.Symbol, output.HeadMove)
		}
	}
}

func (transitionTable TTF) GetFinishState() uint16 {
	// in this program,
	// it assumes state of Turing Machine start by 0
	// and finished by max value of states.
	var maxState uint16 = 0
	for _, outputMap := range transitionTable {
		for _, output := range outputMap {
			if output.State > maxState {
				maxState = output.State
			}
		}
	}
	return maxState
}
