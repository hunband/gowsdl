package xml

import (
	"encoding/xml"
	"github.com/senseyeio/duration"
)

type Duration struct {
	duration.Duration
}

// MarshalXML satisfies xml.Marshaler.
func (d Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(d.String())
}

// UnmarshalXML satisfies xml.Unmarshaler.
func (result *Duration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoded, err := decode(d, start)
	if err != nil {
		return err
	}

	tmp, err := duration.ParseISO8601(decoded)
	if err != nil {
		return err
	}
	result = &Duration{tmp}

	return nil
}
