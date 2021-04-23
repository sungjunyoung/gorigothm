package frog_jmp

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		X      int
		Y      int
		D      int
		expect int
	}{
		{
			X:      10,
			Y:      85,
			D:      30,
			expect: 3,
		},
	}

	for _, test := range tests {
		if Solution(test.X, test.Y, test.D) != test.expect {
			t.Fatal("failed")
		}
	}
}
