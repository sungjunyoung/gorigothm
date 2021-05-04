package remove_duplicates_from_sorted_array

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums       []int
		expect     int
		expectNums []int
	}{
		{
			nums:       []int{0},
			expect:     1,
			expectNums: []int{0},
		},
		{
			nums:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect:     5,
			expectNums: []int{0, 1, 2, 3, 4},
		},
		{
			nums:       []int{1, 1, 2},
			expect:     2,
			expectNums: []int{1, 2},
		},
	}

	for i, test := range tests {
		result := removeDuplicates(test.nums)
		if result != test.expect {
			t.Errorf("test %d failed", i)
		}

		for i := 0; i < result; i++ {
			if test.nums[i] != test.expectNums[i] {
				t.Errorf("test nums %d failed", i)
			}
		}
	}
}
