package compare_the_triplets

import (
	"reflect"
	"testing"
)

func Test_compareTriplets(t *testing.T) {
	tests := []struct {
		a      []int32
		b      []int32
		expect []int32
	}{
		{
			a:      []int32{17, 28, 30},
			b:      []int32{99, 16, 8},
			expect: []int32{2, 1},
		},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(compareTriplets(test.a, test.b), test.expect) {
			t.Errorf("test %d failed", i)
		}
	}
}
