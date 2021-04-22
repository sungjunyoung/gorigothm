package fizz_buzz

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		n      int
		expect []string
	}{
		{
			n:      3,
			expect: []string{"1", "2", "Fizz"},
		},
		{
			n:      15,
			expect: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(fizzBuzz(test.n), test.expect) {
			t.Fatal("failed")
		}
	}
}
