package tm_client

import (
	pb "../proto"
	"encoding/xml"
	"log"
)

// construct transition table
func ReadTtfFromFile(xmlFileName string) *pb.Config {
	return readTtfFromString(stringFromXmlFile(xmlFileName))
}

// construct transition table
func readTtfFromString(xmlString string) *pb.Config {
	// log.Printf("xmlString: %s\n", xmlString)
	return NewConfig(xmlString)
}

func NewConfig(xmlString string) *pb.Config {
	tts := new(pb.Config) // transition table struct
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), tts); err != nil {
		log.Fatalf("Error: TransitionTable XML Unmarshal error: %v\n", err)
	}
	return tts
}
