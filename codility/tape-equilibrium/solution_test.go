package tape_equilibrium

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{3, 1, 2, 4, 3},
			expect: 1,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
