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

var TransitionTable = make(map[uint16]map[string]Output)

func ReadDataFile(xmlFileName string) string {
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
	var lines = make([]string, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}

	// // print lines
	// fmt.Println("## lines")
	// for i, line := range lines {
	// 	fmt.Printf("%2d: %s\n", i + 1, line)
	// }

	// join strings
	return strings.Join(lines[:], "\n")
}

func ParseTransitionTableString(xmlString string) Config {
	// unmarshal (parse)
	var data = Config{}
	// xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), &data); err != nil {
		fmt.Println("!! Error: TransitionTable XML Unmarshal error: ", err)
	}
	return data
}

func PrintTransitionTableXmlByStruct(xmlStruct Config) {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(xmlStruct, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}

func CreateTransitionTable(config Config) {
	var deltaList = config.TuringMachine.TransitionFunction.DeltaList

	for _, delta := range deltaList {
		var input = delta.Input
		var output = delta.Output
		if TransitionTable[input.State] == nil {
			TransitionTable[input.State] = make(map[string]Output)
		}
		TransitionTable[input.State][input.Symbol] = output
	}

	fmt.Printf("input        | output\n")
	fmt.Printf("state symbol | state symbol headmove\n")
	for inputState, outputMap := range TransitionTable {
		for inputSymbol, output := range outputMap {
			fmt.Printf("%5d %6s | %5d %6s %8s\n", inputState, inputSymbol, output.State, output.Symbol, output.HeadMove)
		}
	}
}

func GetFinishState() uint16 {
	// in this program,
	// it assumes state of Turing Machine start by 0
	// and finished by max value of states.
	var maxState uint16 = 0
	for _, outputMap := range TransitionTable {
		for _, output := range outputMap {
			if output.State > maxState {
				maxState = output.State
			}
		}
	}
	return maxState
}
