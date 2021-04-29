package fish

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		B      []int
		expect int
	}{
		{
			A:      []int{4},
			B:      []int{1},
			expect: 1,
		},
		{
			A:      []int{4, 2},
			B:      []int{0, 0},
			expect: 2,
		},
		{
			A:      []int{4, 3, 2, 1, 5},
			B:      []int{0, 1, 0, 0, 0},
			expect: 2,
		},
	}

	for _, test := range tests {
		if Solution(test.A, test.B) != test.expect {
			t.Fatal("failed")
		}
	}
}
