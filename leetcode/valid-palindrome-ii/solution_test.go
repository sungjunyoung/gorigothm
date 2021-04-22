package valid_palindrome_ii

import "testing"

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect bool
	}{
		{
			s:      "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			expect: true,
		},
		{
			s:      "aba",
			expect: true,
		},
		{
			s:      "abca",
			expect: true,
		},
		{
			s:      "abxcddcba",
			expect: true,
		},
		{
			s:      "abxxcddcba",
			expect: false,
		},
	}

	for _, test := range tests {
		if validPalindrome(test.s) != test.expect {
			t.Fatal("failed")
		}
	}
}
