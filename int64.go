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
	_ sql.Scanner              = (*Int64)(nil)
	_ driver.Valuer            = (*Int64)(nil)
	_ json.Marshaler           = (*Int64)(nil)
	_ json.Unmarshaler         = (*Int64)(nil)
	_ encoding.TextMarshaler   = (*Int64)(nil)
	_ encoding.TextUnmarshaler = (*Int64)(nil)
	_ xml.Marshaler            = (*Int64)(nil)
	_ xml.Unmarshaler          = (*Int64)(nil)
)

// Int64 represents a int64 that may be null.
type Int64 struct {
	sql.NullInt64
}

// NewInt64 creates a new Int64.
func NewInt64(i int64) Int64 {
	return Int64{sql.NullInt64{Int64: i, Valid: true}}
}

// NewInt64Ptr creates a new Int64 pointer.
func NewInt64Ptr(i int64) *Int64 {
	v := NewInt64(i)
	return &v
}

// MarshalJSON implements the json.Marshaler interface.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte(jsonNull), nil
	}
	return json.Marshal(i.Int64)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == jsonNull {
		i.Int64, i.Valid = 0, false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int64)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (i Int64) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(i.Int64, 10)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (i *Int64) UnmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		i.Int64, i.Valid = 0, false
		return nil
	}
	i.Valid = true
	i.Int64, err = strconv.ParseInt(string(data), 10, 64)
	return
}

// MarshalXML implements the xml.Marshaler interface.
func (i Int64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !i.Valid {
		start.Attr = []xml.Attr{xsiNilAttr}
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
		return nil
	}
	return e.EncodeElement(i.Int64, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (i *Int64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if isXsiNilAttr(attr) {
			i.Int64, i.Valid = 0, false
			return d.Skip()
		}
	}
	i.Valid = true
	return d.DecodeElement(&i.Int64, &start)
}
