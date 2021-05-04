package factorial_trailing_zeroes

import "testing"

func Test_trailingZeros(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{
			n:      30,
			expect: 7,
		},
		{
			n:      7,
			expect: 1,
		},
		{
			n:      3,
			expect: 0,
		},
		{
			n:      5,
			expect: 1,
		},
		{
			n:      0,
			expect: 0,
		},
	}

	for i, test := range tests {
		if trailingZeroes(test.n) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
