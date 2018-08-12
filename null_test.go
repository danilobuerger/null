// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"encoding"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

const (
	testTypeJson testType = "json"
	testTypeText testType = "text"
	testTypeXml  testType = "xml"
)

type testType string

type nullTestValue interface {
	json.Marshaler
	json.Unmarshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	xml.Marshaler
	xml.Unmarshaler
}

type nullTestValueXmlWrapper struct {
	Namespace string      `xml:"xmlns:xsi,attr"`
	Value     interface{} `xml:"Value"` // *nullTestValue
}

type nullTest struct {
	Type        testType
	New         func() interface{} // Generator of *nullTestValue (and their wrappers)
	Unmarshaled interface{}
	Marshaled   string
}

func nullTestMarshalJSON(t *testing.T, tt nullTest) {
	got, err := json.Marshal(tt.Unmarshaled)
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.Marshaled {
		t.Errorf("got %s; expected %s", got, tt.Marshaled)
	}
}

func nullTestUnmarshalJSON(t *testing.T, tt nullTest) {
	got := tt.New()
	if err := json.Unmarshal([]byte(tt.Marshaled), got); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.Unmarshaled) {
		t.Errorf("got %v; expected %v", got, tt.Unmarshaled)
	}
}

func nullTestMarshalText(t *testing.T, tt nullTest) {
	got, err := tt.Unmarshaled.(nullTestValue).MarshalText()
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.Marshaled {
		t.Errorf("got %s; expected %s", got, tt.Marshaled)
	}
}

func nullTestUnmarshalText(t *testing.T, tt nullTest) {
	got := tt.New().(nullTestValue)
	if err := got.UnmarshalText([]byte(tt.Marshaled)); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.Unmarshaled) {
		t.Errorf("got %v; expected %v", got, tt.Unmarshaled)
	}
}

func nullTestMarshalXML(t *testing.T, tt nullTest) {
	got, err := xml.Marshal(tt.Unmarshaled)
	if err != nil {
		t.Error(err)
	}
	if string(got) != tt.Marshaled {
		t.Errorf("got %s; expected %s", got, tt.Marshaled)
	}
}

func nullTestUnmarshalXML(t *testing.T, tt nullTest) {
	got := tt.New()
	if err := xml.Unmarshal([]byte(tt.Marshaled), got); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, tt.Unmarshaled) {
		t.Errorf("got %v; expected %v", got, tt.Unmarshaled)
	}
}

func nullTestRun(t *testing.T, tests []nullTest) {
	for _, tt := range tests {
		switch tt.Type {
		case testTypeJson:
			t.Run("MarshalJSON", func(t *testing.T) {
				nullTestMarshalJSON(t, tt)
			})
			t.Run("UnmarshalJSON", func(t *testing.T) {
				nullTestUnmarshalJSON(t, tt)
			})
		case testTypeText:
			t.Run("MarshalText", func(t *testing.T) {
				nullTestMarshalText(t, tt)
			})
			t.Run("UnmarshalText", func(t *testing.T) {
				nullTestUnmarshalText(t, tt)
			})
		case testTypeXml:
			t.Run("MarshalXML", func(t *testing.T) {
				nullTestMarshalXML(t, tt)
			})
			t.Run("UnmarshalXML", func(t *testing.T) {
				nullTestUnmarshalXML(t, tt)
			})
		default:
			t.Errorf("Unsupported testType: %#v", tt.Type)
		}
	}
}
