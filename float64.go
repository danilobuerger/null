// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"encoding/xml"
	"strconv"
)

var (
	_ sql.Scanner              = (*Float64)(nil)
	_ driver.Valuer            = (*Float64)(nil)
	_ json.Marshaler           = (*Float64)(nil)
	_ json.Unmarshaler         = (*Float64)(nil)
	_ encoding.TextMarshaler   = (*Float64)(nil)
	_ encoding.TextUnmarshaler = (*Float64)(nil)
	_ xml.Marshaler            = (*Float64)(nil)
	_ xml.Unmarshaler          = (*Float64)(nil)
)

// Float64 represents a float64 that may be null.
type Float64 struct {
	sql.NullFloat64
}

// NewFloat64 creates a new Float64.
func NewFloat64(f float64) Float64 {
	return Float64{sql.NullFloat64{Float64: f, Valid: true}}
}

// NewFloat64Ptr creates a new Float64 pointer.
func NewFloat64Ptr(f float64) *Float64 {
	v := NewFloat64(f)
	return &v
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte(jsonNull), nil
	}
	return json.Marshal(f.Float64)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == jsonNull {
		f.Float64, f.Valid = 0, false
		return nil
	}
	f.Valid = true
	return json.Unmarshal(data, &f.Float64)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (f Float64) MarshalText() ([]byte, error) {
	if !f.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatFloat(f.Float64, 'f', -1, 64)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (f *Float64) UnmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		f.Float64, f.Valid = 0, false
		return nil
	}
	f.Valid = true
	f.Float64, err = strconv.ParseFloat(string(data), 64)
	return
}

// MarshalXML implements the xml.Marshaler interface.
func (f Float64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !f.Valid {
		start.Attr = []xml.Attr{xsiNilAttr}
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
		return nil
	}
	return e.EncodeElement(f.Float64, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (f *Float64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if isXsiNilAttr(attr) {
			f.Float64, f.Valid = 0, false
			return d.Skip()
		}
	}
	f.Valid = true
	return d.DecodeElement(&f.Float64, &start)
}
