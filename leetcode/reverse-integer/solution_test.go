package reverse_integer

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		x      int
		expect int
	}{
		{
			x:      123,
			expect: 321,
		},
		{
			x:      1534236469,
			expect: 0,
		},
		{
			x:      -123,
			expect: -321,
		},
		{
			x:      120,
			expect: 21,
		},
		{
			x:      0,
			expect: 0,
		},
	}

	for _, test := range tests {
		result := reverse(test.x)
		if !reflect.DeepEqual(result, test.expect) {
			t.Fatal("failed")
		}
	}
}
