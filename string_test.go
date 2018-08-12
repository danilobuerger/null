// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import "testing"

func stringNullValue() interface{} {
	return &String{}
}

func newStringNullValueXmlWrapper() interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: stringNullValue()}
}

func newStringValueXmlWrapper(value *String) interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: value}
}

func TestString(t *testing.T) {
	tests := []nullTest{
		// null values
		{testTypeJson, stringNullValue, stringNullValue(), `null`},
		{testTypeText, stringNullValue, NewStringPtr(``), ``},
		{testTypeXml, stringNullValue, stringNullValue(), `<String xsi:nil="true"></String>`},
		{
			testTypeXml,
			newStringNullValueXmlWrapper,
			newStringNullValueXmlWrapper(),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value xsi:nil="true"></Value></nullTestValueXmlWrapper>`,
		},
		// `` values
		{testTypeJson, stringNullValue, NewStringPtr(``), `""`},
		{testTypeText, stringNullValue, NewStringPtr(``), ``},
		{testTypeXml, stringNullValue, NewStringPtr(``), `<String></String>`},
		{
			testTypeXml,
			newStringNullValueXmlWrapper,
			newStringValueXmlWrapper(NewStringPtr(``)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value></Value></nullTestValueXmlWrapper>`,
		},
		// `foo` values
		{testTypeJson, stringNullValue, NewStringPtr(`foo`), `"foo"`},
		{testTypeText, stringNullValue, NewStringPtr(`foo`), `foo`},
		{testTypeXml, stringNullValue, NewStringPtr(`foo`), `<String>foo</String>`},
		{
			testTypeXml,
			newStringNullValueXmlWrapper,
			newStringValueXmlWrapper(NewStringPtr(`foo`)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value>foo</Value></nullTestValueXmlWrapper>`,
		},
	}

	nullTestRun(t, tests)
}
