package perm_missing_elem

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{2, 3, 1, 5},
			expect: 4,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
