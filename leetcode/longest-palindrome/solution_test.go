package longest_palindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{
			s:      "abccccdd",
			expect: 7,
		},
		{
			s:      "a",
			expect: 1,
		},
		{
			s:      "bb",
			expect: 2,
		},
	}

	for i, test := range tests {
		if longestPalindrome(test.s) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
