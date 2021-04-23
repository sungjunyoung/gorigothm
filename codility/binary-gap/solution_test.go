package binary_gap

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		N      int
		expect int
	}{
		{
			N:      9,
			expect: 2,
		},
		{
			N:      529,
			expect: 4,
		},
		{
			N:      20,
			expect: 1,
		},
		{
			N:      15,
			expect: 0,
		},
		{
			N:      32,
			expect: 0,
		},
	}

	for _, test := range tests {
		if Solution(test.N) != test.expect {
			t.Fatal("failed")
		}
	}
}
