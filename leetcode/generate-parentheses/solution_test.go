package generate_parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		n      int
		expect []string
	}{
		{n: 3, expect: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{n: 1, expect: []string{"()"}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(generateParenthesis(test.n), test.expect) {
			t.Fatal("failed")
		}
	}
}
