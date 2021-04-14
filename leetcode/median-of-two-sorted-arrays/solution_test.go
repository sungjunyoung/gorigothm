package median_of_two_sorted_arrays

import (
	"reflect"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{
			nums1:  []int{1, 3},
			nums2:  []int{2},
			expect: 2,
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			expect: 2.5,
		},
		{
			nums1:  []int{0, 0},
			nums2:  []int{0, 0},
			expect: 0,
		},
		{
			nums1:  []int{},
			nums2:  []int{1},
			expect: 1,
		},
		{
			nums1:  []int{2},
			nums2:  []int{},
			expect: 2,
		},
	}

	for _, test := range tests {
		result := findMedianSortedArrays(test.nums1, test.nums2)
		if !reflect.DeepEqual(result, test.expect) {
			t.Fatal("failed")
		}
	}
}
