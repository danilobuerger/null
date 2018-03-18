// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"encoding/xml"
)

var (
	_ sql.Scanner              = (*String)(nil)
	_ driver.Valuer            = (*String)(nil)
	_ json.Marshaler           = (*String)(nil)
	_ json.Unmarshaler         = (*String)(nil)
	_ encoding.TextMarshaler   = (*String)(nil)
	_ encoding.TextUnmarshaler = (*String)(nil)
	_ xml.Marshaler            = (*String)(nil)
	_ xml.Unmarshaler          = (*String)(nil)
)

// String represents a string that may be null.
type String struct {
	sql.NullString
}

// NewString creates a new String.
func NewString(s string) String {
	return String{sql.NullString{String: s, Valid: true}}
}

// NewStringPtr creates a new String pointer.
func NewStringPtr(s string) *String {
	v := NewString(s)
	return &v
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte(jsonNull), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == jsonNull {
		s.String, s.Valid = "", false
		return nil
	}
	s.Valid = true
	return json.Unmarshal(data, &s.String)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s String) MarshalText() ([]byte, error) {
	if !s.Valid {
		return []byte{}, nil
	}
	return []byte(s.String), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *String) UnmarshalText(data []byte) (err error) {
	s.String, s.Valid = string(data), true
	return
}

// MarshalXML implements the xml.Marshaler interface.
func (s String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !s.Valid {
		start.Attr = []xml.Attr{xsiNilAttr}
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
		return nil
	}
	return e.EncodeElement(s.String, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (s *String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if isXsiNilAttr(attr) {
			s.String, s.Valid = "", false
			return d.Skip()
		}
	}
	s.Valid = true
	return d.DecodeElement(&s.String, &start)
}
