package intersection_of_two_arrays_ii

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{
			nums1:  []int{1, 2, 2, 1},
			nums2:  []int{2, 2},
			expect: []int{2, 2},
		},
	}

	for _, test := range tests {
		result := intersect(test.nums1, test.nums2)
		if len(result) != len(test.expect) {
			t.Errorf("failed")
		}
		sort.Ints(result)
		sort.Ints(test.expect)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("failed")
		}
	}
}
