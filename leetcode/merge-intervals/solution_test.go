package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expect    [][]int
	}{
		{
			intervals: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			expect: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},

		{
			intervals: [][]int{
				{1, 4}, {4, 5},
			},
			expect: [][]int{
				{1, 5},
			},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(merge(test.intervals), test.expect) {
			t.Fatal("failed")
		}
	}
}
