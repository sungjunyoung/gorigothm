package plus_one

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		digits []int
		expect []int
	}{
		{
			digits: []int{9, 9, 9},
			expect: []int{1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(plusOne(test.digits), test.expect) {
			t.Error("failed")
		}
	}
}
