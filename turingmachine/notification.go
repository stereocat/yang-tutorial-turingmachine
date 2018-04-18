package turingmachine

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Notification struct {
	XMLName   xml.Name `xml:"notification"`
	EventTime string   `xml:"eventTime"`
	Xmlns     string   `xml:"xmlns,attr"`
	Halted    Halted   `xml:"halted"`
}
type Halted struct {
	XMLName xml.Name `xml:"halted"`
	Xmlns   string   `xml:"xmlns,attr"`
	State   uint16   `xml:"state"`
}

func newNotification(state uint16) *Notification {
	return &Notification{
		EventTime: time.Now().String(),
		Xmlns:     "urn:ietf:params:xml:ns:netconf:notification:1.0",
		Halted: Halted{
			State: state,
			Xmlns: "http://example.net/turing-machine",
		}}
}

func (notify *Notification) PrintXml() {
	var xmlBuf, err = xml.MarshalIndent(notify, "", "  ")
	if err != nil {
		fmt.Println("!! Error: XML Marshal err: ", err)
	}
	fmt.Println(string(xmlBuf))
}
