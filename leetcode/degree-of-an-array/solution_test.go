package degree_of_an_array

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			expect: 7,
		},
		{
			nums:   []int{1, 2, 2, 3, 1, 4, 2},
			expect: 6,
		},
		{
			nums:   []int{1, 2, 2, 3, 1},
			expect: 2,
		},
	}

	for _, test := range tests {
		if findShortestSubArray(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
