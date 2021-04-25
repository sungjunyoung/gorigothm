package perm_check

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{4, 1, 3, 2},
			expect: 1,
		},
		{
			A:      []int{4, 1, 3},
			expect: 0,
		},
		{
			A:      []int{2},
			expect: 0,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
