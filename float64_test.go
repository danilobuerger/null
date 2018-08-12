// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"testing"
)

func float64NullValue() interface{} {
	return &Float64{}
}

func newFloat64NullValueXmlWrapper() interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: float64NullValue()}
}

func newFloat64ValueXmlWrapper(value *Float64) interface{} {
	return &nullTestValueXmlWrapper{Namespace: xsiNamespace, Value: value}
}

func TestFloat64(t *testing.T) {
	tests := []nullTest{
		// null values
		{testTypeJson, float64NullValue, float64NullValue(), `null`},
		{testTypeText, float64NullValue, float64NullValue(), ``},
		{testTypeXml, float64NullValue, float64NullValue(), `<Float64 xsi:nil="true"></Float64>`},
		{
			testTypeXml,
			newFloat64NullValueXmlWrapper,
			newFloat64NullValueXmlWrapper(),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value xsi:nil="true"></Value></nullTestValueXmlWrapper>`,
		},
		// 0.0 values
		{testTypeJson, float64NullValue, NewFloat64Ptr(0.0), `0`},
		{testTypeText, float64NullValue, NewFloat64Ptr(0.0), `0`},
		{testTypeXml, float64NullValue, NewFloat64Ptr(0.0), `<Float64>0</Float64>`},
		{
			testTypeXml,
			newFloat64NullValueXmlWrapper,
			newFloat64ValueXmlWrapper(NewFloat64Ptr(0.0)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value>0</Value></nullTestValueXmlWrapper>`,
		},
		// 1.12 values
		{testTypeJson, float64NullValue, NewFloat64Ptr(1.12), `1.12`},
		{testTypeText, float64NullValue, NewFloat64Ptr(1.12), `1.12`},
		{testTypeXml, float64NullValue, NewFloat64Ptr(1.12), `<Float64>1.12</Float64>`},
		{
			testTypeXml,
			newFloat64NullValueXmlWrapper,
			newFloat64ValueXmlWrapper(NewFloat64Ptr(1.12)),
			`<nullTestValueXmlWrapper xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Value>1.12</Value></nullTestValueXmlWrapper>`,
		},
	}

	nullTestRun(t, tests)
}
