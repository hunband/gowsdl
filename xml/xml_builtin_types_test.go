package xml

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

type TestDateTime struct {
	XMLName  xml.Name `xml:"testDateTime"`
	SomeDate DateTime `xml:"dateToTest"`
}

type TestDateTimeAttr struct {
	XMLName      xml.Name `xml:"testDateTime"`
	SomeDateAttr DateTime `xml:"dateToTest,attr"`
}

type TestDuration struct {
	XMLName      xml.Name      `xml:"testDuration"`
	SomeDuration time.Duration `xml:"durationToTest"`
}

func TestDateTime_UnmarshalXML(t *testing.T) {
	var dateTimeAsFieldXml = []byte(`<testDateTime><dateToTest>2002-05-30T09:30:10.555+23:00</dateToTest></testDateTime>`)
	//var dateTimeAsAttrXml = []byte(`<testDateTime dateToTest=20140729></testDateTime>`)

	var sample TestDateTime
	xml.Unmarshal(dateTimeAsFieldXml, &sample)
	fmt.Printf("enterdate field: %s\n", sample)

	//layout := "2006-01-02T15:04:05.000Z"
	//str := "2014-11-12T11:45:26.371Z"
	//tm, err := time.Parse(layout, str)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(tm)
}

func TestDuration_UnmarshalXML(t *testing.T) {

	//var durationAsFieldXml = []byte(`<testDuration><durationToTest>P1Y2MT2H</durationToTest></testDuration>`)
	////var dateTimeAsAttrXml = []byte(`<testDateTime dateToTest=20140729></testDateTime>`)
	//
	//var sample TestDuration
	//err := xml.Unmarshal(durationAsFieldXml, &sample)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("enterdate field: %s\n", sample)

	d := time.Duration(1)
	fmt.Println(d * time.Hour * 24)
	//layout := "2006-01-02T15:04:05.000Z"
	//str := "2014-11-12T11:45:26.371Z"
	//tm, err := time.Parse(layout, str)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(tm)
}
