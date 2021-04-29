package nesting

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		S      string
		expect int
	}{
		{
			S:      "(()",
			expect: 0,
		},
		{
			S:      "(()(())())",
			expect: 1,
		},
		{
			S:      "())",
			expect: 0,
		},
	}

	for _, test := range tests {
		if Solution(test.S) != test.expect {
			t.Fatal("failed")
		}
	}
}
