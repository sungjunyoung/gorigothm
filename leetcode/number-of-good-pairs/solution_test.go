package number_of_good_pairs

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 2, 3, 1, 1, 3},
			expect: 4,
		},
		{
			nums:   []int{1, 1, 1, 1},
			expect: 6,
		},
	}

	for _, test := range tests {
		if numIdenticalPairs(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
