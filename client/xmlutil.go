package tm_client

import (
	pb "../proto"
	"encoding/xml"
	"log"
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

func NewConfig(xmlString string) *pb.Config {
	tts := new(pb.Config) // transition table struct
	// unmarshal (parse); xml.Unmarshal arg must be []byte
	if err := xml.Unmarshal([]byte(xmlString), tts); err != nil {
		log.Fatalf("Error: TransitionTable XML Unmarshal error: ", err)
	}
	return tts
}

func TMXmlString(tm *pb.TuringMachine) string {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tm, "", "  ")
	if err != nil {
		log.Fatalf("Error: XML Marshal err: ", err)
	}
	return string(xmlBuf)
}
