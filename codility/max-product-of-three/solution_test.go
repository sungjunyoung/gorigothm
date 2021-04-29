package max_product_of_three

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{-3, 1, 2, -2, 5, 6},
			expect: 60,
		},
		{
			A:      []int{-3, -2, -1, 0, 1, 2},
			expect: 12,
		},
		{
			A:      []int{-3, -2, -1, 0, 0, 6},
			expect: 36,
		},
		{
			A:      []int{-3, -2, -1, 0, 0, 0},
			expect: 0,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
