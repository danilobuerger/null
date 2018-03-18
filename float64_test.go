// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"testing"
)

func float64NullValue() nullTestValue {
	return &Float64{}
}

func TestFloat64(t *testing.T) {
	tests := []nullTest{
		{float64NullValue, float64NullValue(), `null`, float64NullValue(), ``, float64NullValue(), `<Float64 xsi:nil="true"></Float64>`},
		{float64NullValue, NewFloat64Ptr(0.0), `0`, NewFloat64Ptr(0.0), `0`, NewFloat64Ptr(0.0), `<Float64>0</Float64>`},
		{float64NullValue, NewFloat64Ptr(1.12), `1.12`, NewFloat64Ptr(1.12), `1.12`, NewFloat64Ptr(1.12), `<Float64>1.12</Float64>`},
	}

	nullTestRun(t, tests)
}
