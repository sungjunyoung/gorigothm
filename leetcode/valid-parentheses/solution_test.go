package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		s      string
		expect bool
	}{
		{
			s:      "()",
			expect: true,
		},
		{
			s:      "()[]{}",
			expect: true,
		},

		{
			s:      "([)]",
			expect: false,
		},
	}

	for _, test := range tests {
		if isValid(test.s) != test.expect {
			t.Fatal("failed")
		}
	}
}
