package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnySimpleType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string
type Notation string

type NormalizedString string
type Token string
type Language string
type ID string
type IDREF string
type IDREFS []string
type ENTITY string
type ENTITIES []string
type NMTOKEN string
type NMTOKENS []string
type Name string
type NCName string

type DateTime struct {
	time.Time
}

type Time struct {
	time.Time
}

type Date struct {
	time.Time
}

type GYearMonth struct {
	time.Time
}

type GYear struct {
	time.Time
}

type GMonthDay struct {
	time.Time
}

type GDay struct {
	time.Time
}

type GMonth struct {
	time.Time
}

var validDateTimeLayouts = []string{
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05.99",
	"2006-01-02T15:04:05.99Z",
	"2006-01-02T15:04:05.99-07:00",
}

var validTimeLayouts = []string{
	"15:04:05",
	"15:04:05-07:00",
	"15:04:05.99",
	"15:04:05.99Z",
	"15:04:05.99-07:00",
}

var validDateLayouts = []string{
	"2006-01-02",
	"2006-01-02Z",
	"2006-01-02-07:00",
}

var validGYearMonthLayouts = []string{
	"2006-01",
	"2006-01Z",
	"2006-01-07:00",
}

var validGYearLayouts = []string{
	"2006",
	"2006Z",
	"2006-07:00",
}

var validGMonthDayLayouts = []string{
	"--01-02",
	"--01-02Z",
	"--01-02-07:00",
}

var validGDayLayouts = []string{
	"---02",
	"---02Z",
	"---02-07:00",
}

var validGMonthLayouts = []string{
	"--01",
	"--01Z",
	"--01-07:00",
}

var typeLayouts = map[string][]string{
	"DateTime":   validDateTimeLayouts,
	"Time":       validTimeLayouts,
	"Date":       validDateLayouts,
	"gYearMonth": validGYearMonthLayouts,
	"gYear":      validGYearLayouts,
	"gMonthDay":  validGMonthDayLayouts,
	"gDay":       validGDayLayouts,
	"gMonth":     validGMonthLayouts,
}

func (result *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("DateTime", decoded, result)
}

func (result *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("Time", decoded, result)
}

func (result *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("Date", decoded, result)
}

func (result *GYearMonth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("gYearMonth", decoded, result)
}

func (result *GYear) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("gYear", decoded, result)
}

func (result *GMonthDay) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("gMonthDay", decoded, result)
}

func (result *GDay) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("gDay", decoded, result)
}

func (result *GMonth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}
	return tryToParseTime("gMonth", decoded, result)
}

//func (result *Duration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	decoded, err := decode(d, start)
//	if err != nil {
//		return err
//	}
//
//}

func decode(d *xml.Decoder, start xml.StartElement) (string, error) {
	var timeString string
	err := d.DecodeElement(&timeString, &start)
	return timeString, err
}

func tryToParseTime(_type string, decodedXsdTime string, result interface{}) error {
	for _, layout := range typeLayouts[_type] {
		parse, parseError := time.Parse(layout, decodedXsdTime)
		if parseError == nil {
			result = &DateTime{parse}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Cannot decode %s to %s type", decodedXsdTime, _type))
}

//
//func tryToParseDuration(_type string, decodedXsdDuration string, result interface{}) error {
//	decodedXsdDuration
//	return errors.New(fmt.Sprintf("Cannot decode %s to %s type", decodedXsdTime, _type))
//}
