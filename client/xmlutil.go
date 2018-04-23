package tm_client

import (
	pb "../proto"
	"bufio"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

// func (c *Config) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	uu := struct {
// 		XMLName       xml.Name      `xml:"config"`
// 		xmlns         string        `xml:"xmlns,attr"`
// 		TuringMachine TuringMachine `xml:"turing-machine"`
// 	}{}
// 	if err := d.DecodeElement(&uu, &start); err != nil {
// 		return err
// 	}
// 	// log.Printf("B decoder: %v\nstartelement: %v\nConfig: %v\nuu: %v\n", d, start, c, uu)
//     log.Printf("uu: %v\n", uu)
// 	*c = Config{TuringMachine: &uu.TuringMachine}
// 	return nil
// }

func stringFromXmlFile(xmlFileName string) string {
	xmlFile, err := os.Open(xmlFileName)
	if err != nil {
		log.Fatalf("Error: Cannot read file:%s\n", xmlFileName)
		os.Exit(1)
	} else {
		defer xmlFile.Close() // close finished readXmlFile
	}
	return openXmlFile(xmlFile)
}

func openXmlFile(xmlFile *os.File) string {
	// read data from file
	var scanner = bufio.NewScanner(xmlFile)
	var lines = make([]string, 0) // multiple lines
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n") // convert to single line
}

func TMXmlString(tm *pb.TuringMachine) string {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tm, "", "  ")
	if err != nil {
		log.Fatalf("Error: XML Marshal err: ", err)
	}
	return string(xmlBuf)
}
