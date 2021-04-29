package passing_cars

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{0, 1, 0, 1, 1},
			expect: 5,
		},
		{
			A:      []int{1, 1, 1, 0, 1, 0, 0},
			expect: 1,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
