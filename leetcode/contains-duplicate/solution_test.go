package contains_duplicate

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{
			nums:   []int{1, 2, 3, 1},
			expect: true,
		},
		{
			nums:   []int{1, 2, 3, 4},
			expect: false,
		},
		{
			nums:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expect: true,
		},
	}

	for i, test := range tests {
		if containsDuplicate(test.nums) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
