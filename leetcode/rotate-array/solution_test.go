package rotate_array

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			expect: []int{3, 99, -1, -100},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		rotate(test.nums, test.k)
		if !reflect.DeepEqual(test.nums, test.expect) {
			t.Errorf("test %d failed", i)
		}
	}
}
