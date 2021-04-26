package hamming_distance

import "testing"

func Test_hammingDistance(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		expect int
	}{
		{
			x:      1,
			y:      4,
			expect: 2,
		},
	}

	for _, test := range tests {
		if hammingDistance(test.x, test.y) != test.expect {
			t.Fatal("failed")
		}
	}
}
