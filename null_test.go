// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"encoding"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

type nullTestValue interface {
	json.Marshaler
	json.Unmarshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	xml.Marshaler
	xml.Unmarshaler
}

type nullTest struct {
	new       func() nullTestValue
	jsonValue nullTestValue
	json      string
	textValue nullTestValue
	text      string
	xmlValue  nullTestValue
	xml       string
}

func nullTestMarshalJSON(t *testing.T, tt nullTest) {
	got, err := json.Marshal(tt.jsonValue)
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.json {
		t.Errorf("got %s; expected %s", got, tt.json)
	}
}

func nullTestUnmarshalJSON(t *testing.T, tt nullTest) {
	got := tt.new()
	if err := json.Unmarshal([]byte(tt.json), got); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.jsonValue) {
		t.Errorf("got %v; expected %v", got, tt.jsonValue)
	}
}

func nullTestMarshalText(t *testing.T, tt nullTest) {
	got, err := tt.textValue.MarshalText()
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.text {
		t.Errorf("got %s; expected %s", got, tt.text)
	}
}

func nullTestUnmarshalText(t *testing.T, tt nullTest) {
	got := tt.new()
	if err := got.UnmarshalText([]byte(tt.text)); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.textValue) {
		t.Errorf("got %v; expected %v", got, tt.textValue)
	}
}

func nullTestMarshalXML(t *testing.T, tt nullTest) {
	got, err := xml.Marshal(tt.xmlValue)
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.xml {
		t.Errorf("got %s; expected %s", got, tt.xml)
	}
}

func nullTestUnmarshalXML(t *testing.T, tt nullTest) {
	got := tt.new()
	if err := xml.Unmarshal([]byte(tt.xml), got); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.xmlValue) {
		t.Errorf("got %v; expected %v", got, tt.xmlValue)
	}
}

func nullTestRun(t *testing.T, tests []nullTest) {
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			nullTestMarshalJSON(t, tt)
			nullTestUnmarshalJSON(t, tt)
			nullTestMarshalText(t, tt)
			nullTestUnmarshalText(t, tt)
			nullTestMarshalXML(t, tt)
			nullTestUnmarshalXML(t, tt)
		})
	}
}
