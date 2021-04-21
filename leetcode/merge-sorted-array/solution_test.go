package merge_sorted_array

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.expect) {
			t.Fatal("failed")
		}
	}
}
