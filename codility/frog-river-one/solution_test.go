package frog_river_one

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		X      int
		A      []int
		expect int
	}{
		{
			X:      5,
			A:      []int{1, 3, 1, 4, 2, 3, 5, 4},
			expect: 6,
		},
	}

	for _, test := range tests {
		if Solution(test.X, test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
