package contains_duplicate_ii

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect bool
	}{
		{
			nums:   []int{1, 2, 3, 1},
			k:      3,
			expect: true,
		},
		{
			nums:   []int{1, 0, 1, 1},
			k:      1,
			expect: true,
		},
		{
			nums:   []int{1, 2, 3, 1, 2, 3},
			k:      2,
			expect: false,
		},
	}

	for i, test := range tests {
		if containsNearbyDuplicate(test.nums, test.k) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
