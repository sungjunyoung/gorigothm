package quicksort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	tests := []struct {
		arr    []int
		expect []int
	}{
		{
			arr:    []int{2, 5, 3, 6, 7, 8, 9, 0, 5, 4, 2, 4, 5},
			expect: []int{0, 2, 2, 3, 4, 4, 5, 5, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		quickSort(test.arr)
		t.Log(test.arr)
		if !reflect.DeepEqual(test.arr, test.expect) {
			t.Error("failed")
		}
	}
}
