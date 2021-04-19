package trapping_rain_water

import "testing"

func Test_trap(t *testing.T) {
	tests := []struct {
		height []int
		expect int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expect: 6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			expect: 9,
		},
	}

	for _, test := range tests {
		if trap(test.height) != test.expect {
			t.Fatal("failed")
		}
	}
}
