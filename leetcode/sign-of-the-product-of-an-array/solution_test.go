package sign_of_the_product_of_an_array

import "testing"

func Test_arraySign(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{-1, -2, -3, -4, 3, 2, 1},
			expect: 1,
		},
		{
			nums:   []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24},
			expect: -1,
		},
	}

	for _, test := range tests {
		if arraySign(test.nums) != test.expect {
			t.Fatal("failed")
		}
	}
}
