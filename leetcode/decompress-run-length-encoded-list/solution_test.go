package decompress_run_length_encoded_list

import (
	"reflect"
	"testing"
)

func Test_decompressRLElist(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			expect: []int{2, 4, 4, 4},
		},
		{
			nums:   []int{1, 1, 2, 3},
			expect: []int{1, 3, 3},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(decompressRLElist(test.nums), test.expect) {
			t.Fatal("failed")
		}
	}
}
