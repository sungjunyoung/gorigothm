package count_div

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      int
		B      int
		K      int
		expect int
	}{
		{
			A:      11,
			B:      345,
			K:      17,
			expect: 20,
		},
		{
			A:      6,
			B:      11,
			K:      2,
			expect: 3,
		},
		{
			A:      3,
			B:      12,
			K:      3,
			expect: 4,
		},
	}

	for _, test := range tests {
		if Solution(test.A, test.B, test.K) != test.expect {
			t.Fatal("failed")
		}
	}
}
