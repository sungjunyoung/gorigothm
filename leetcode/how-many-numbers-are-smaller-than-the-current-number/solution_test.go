package how_many_numbers_are_smaller_than_the_current_number

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{8, 1, 2, 2, 3},
			expect: []int{4, 0, 1, 1, 3},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(smallerNumbersThanCurrent(test.nums), test.expect) {
			t.Fatal("failed")
		}
	}
}
