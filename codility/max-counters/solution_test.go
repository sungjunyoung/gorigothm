package max_counters

import (
	"reflect"
	"testing"
)

func Test_Solution(t *testing.T) {
	tests := []struct {
		N      int
		A      []int
		expect []int
	}{
		{
			N:      5,
			A:      []int{3, 4, 4, 6, 1, 4, 4},
			expect: []int{3, 2, 2, 4, 2},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(Solution(test.N, test.A), test.expect) {
			t.Fatal("failed")
		}
	}
}
