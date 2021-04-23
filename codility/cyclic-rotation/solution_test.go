package cyclic_rotation

import (
	"reflect"
	"testing"
)

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		K      int
		expect []int
	}{
		{
			A:      []int{3, 8, 9, 7, 6},
			K:      3,
			expect: []int{9, 7, 6, 3, 8},
		},
		{
			A:      []int{0, 0, 0},
			K:      1,
			expect: []int{0, 0, 0},
		},
		{
			A:      []int{1, 2, 3, 4},
			K:      4,
			expect: []int{1, 2, 3, 4},
		},
		{
			A:      []int{1, 2, 3, 4},
			K:      2,
			expect: []int{3, 4, 1, 2},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(Solution(test.A, test.K), test.expect) {
			t.Fatal("failed")
		}
	}
}
