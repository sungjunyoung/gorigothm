package reverse_string

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		s      []byte
		expect []byte
	}{
		{
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			expect: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for _, test := range tests {
		reverseString(test.s)
		if !reflect.DeepEqual(test.s, test.expect) {
			t.Fatal("failed")
		}
	}
}
