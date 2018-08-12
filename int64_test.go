// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import "testing"

func int64NullValue() interface{} {
	return &Int64{}
}

func newInt64NullValueXmlWrapper() interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: int64NullValue()}
}

func newInt64ValueXmlWrapper(value *Int64) interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: value}
}

func TestInt64(t *testing.T) {
	tests := []nullTest{
		// null values
		{testTypeJson, int64NullValue, int64NullValue(), `null`},
		{testTypeText, int64NullValue, int64NullValue(), ``},
		{testTypeXml, int64NullValue, int64NullValue(), `<Int64 xsi:nil="true"></Int64>`},
		{
			testTypeXml,
			newInt64NullValueXmlWrapper,
			newInt64NullValueXmlWrapper(),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value xsi:nil="true"></Value></nullTestValueXmlWrapper>`,
		},
		// 0 values
		{testTypeJson, int64NullValue, NewInt64Ptr(0), `0`},
		{testTypeText, int64NullValue, NewInt64Ptr(0), `0`},
		{testTypeXml, int64NullValue, NewInt64Ptr(0), `<Int64>0</Int64>`},
		{
			testTypeXml,
			newInt64NullValueXmlWrapper,
			newInt64ValueXmlWrapper(NewInt64Ptr(0)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value>0</Value></nullTestValueXmlWrapper>`,
		},
		// 1 values
		{testTypeJson, int64NullValue, NewInt64Ptr(1), `1`},
		{testTypeText, int64NullValue, NewInt64Ptr(1), `1`},
		{testTypeXml, int64NullValue, NewInt64Ptr(1), `<Int64>1</Int64>`},
		{
			testTypeXml,
			newInt64NullValueXmlWrapper,
			newInt64ValueXmlWrapper(NewInt64Ptr(1)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value>1</Value></nullTestValueXmlWrapper>`,
		},
	}

	nullTestRun(t, tests)
}
