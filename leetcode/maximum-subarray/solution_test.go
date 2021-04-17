package maximum_subarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{5, 4, -1, 7, 8},
			expect: 23,
		},
		{
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
	}

	for _, test := range tests {
		if maxSubArray(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
