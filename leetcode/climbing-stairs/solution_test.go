package climbing_stairs

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{
			n:      2,
			expect: 2,
		},
		{
			n:      3,
			expect: 3,
		},
	}

	for _, test := range tests {
		if climbStairs(test.n) != test.expect {
			t.Fatal("failed")
		}
	}
}
