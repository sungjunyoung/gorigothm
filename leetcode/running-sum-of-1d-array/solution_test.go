package running_sum_of_1d_array

import (
	"testing"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			expect: []int{1, 3, 6, 10},
		},
		{
			nums:   []int{1, 1, 1, 1, 1},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			nums:   []int{3, 1, 2, 10, 1},
			expect: []int{3, 4, 6, 16, 17},
		},
	}

	for _, test := range tests {
		result := runningSum(test.nums)
		if len(result) != len(test.expect) {
			t.Fatal("length should be equal")
		}

		for i := range result {
			if result[i] != test.expect[i] {
				t.Fatal("failed")
			}
		}
	}
}
