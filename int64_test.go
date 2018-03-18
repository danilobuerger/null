// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import (
	"testing"
)

func int64NullValue() nullTestValue {
	return &Int64{}
}

func TestInt64(t *testing.T) {
	tests := []nullTest{
		{int64NullValue, int64NullValue(), `null`, ``, `<Int64 xsi:nil="true"></Int64>`},
		{int64NullValue, NewInt64Ptr(0), `0`, `0`, `<Int64>0</Int64>`},
		{int64NullValue, NewInt64Ptr(1), `1`, `1`, `<Int64>1</Int64>`},
	}

	nullTestRun(t, tests)
}
