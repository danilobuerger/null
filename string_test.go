// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"testing"
)

func stringNullValue() nullTestValue {
	return &String{}
}

func TestString(t *testing.T) {
	tests := []nullTest{
		{stringNullValue, stringNullValue(), `null`, NewStringPtr(``), ``, stringNullValue(), `<String xsi:nil="true"></String>`},
		{stringNullValue, NewStringPtr(``), `""`, NewStringPtr(``), ``, NewStringPtr(``), `<String></String>`},
		{stringNullValue, NewStringPtr(`foo`), `"foo"`, NewStringPtr(`foo`), `foo`, NewStringPtr(`foo`), `<String>foo</String>`},
	}

	nullTestRun(t, tests)
}
