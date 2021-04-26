package single_number

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{2, 2, 1},
			expect: 1,
		},
		{
			nums:   []int{4, 1, 2, 1, 2},
			expect: 4,
		},
	}

	for _, test := range tests {
		if singleNumber(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
