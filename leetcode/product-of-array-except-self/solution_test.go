package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			expect: []int{24, 12, 8, 6},
		},
		{
			nums:   []int{-1, 1, 0, -3, 3},
			expect: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(productExceptSelf(test.nums), test.expect) {
			t.Fatal("failed")
		}
	}
}
