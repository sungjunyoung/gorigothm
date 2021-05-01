package majority_element

import "testing"

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{3, 2, 3},
			expect: 3,
		},
		{
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			expect: 2,
		},
	}

	for _, test := range tests {
		if majorityElement(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
