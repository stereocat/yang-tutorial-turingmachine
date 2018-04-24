package tmclient

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func stringFromXMLFile(xmlFileName string) string {
	xmlFile, err := os.Open(xmlFileName)
	if err != nil {
		log.Fatalf("Error: Cannot read file:%s\n", xmlFileName)
	} else {
		defer xmlFile.Close() // close finished readXmlFile
	}
	return openXMLFile(xmlFile)
}

func openXMLFile(xmlFile *os.File) string {
	// read data from file
	var scanner = bufio.NewScanner(xmlFile)
	var lines = make([]string, 0) // multiple lines
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return strings.Join(lines[:], "\n") // convert to single line
}
