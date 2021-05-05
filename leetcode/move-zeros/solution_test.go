package move_zeros

import (
	"reflect"
	"testing"
)

func Test_moveZeros(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{0, 1, 0, 3, 12},
			expect: []int{1, 3, 12, 0, 0},
		},
		{
			nums:   []int{0, 1, 0, 0, 2, 3, 0, 4},
			expect: []int{1, 2, 3, 4, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.nums)

		if !reflect.DeepEqual(test.nums, test.expect) {
			t.Errorf("failed")
		}
	}
}
