package missing_integer

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{1, 3, 6, 4, 1, 2},
			expect: 5,
		},
		{
			A:      []int{1, 2, 3},
			expect: 4,
		},
		{
			A:      []int{-1, -3},
			expect: 1,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
