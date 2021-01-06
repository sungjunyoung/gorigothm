package two_sum_test

import (
	"reflect"
	"testing"

	. "github.com/sungjunyoung/gorigothm/leetcode/two_sum"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{1, 0},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{2, 1},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			expect: []int{1, 0},
		},
	}

	for _, test := range tests {
		result := TwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expect) {
			t.Fatal("failed")
		}
	}
}
