package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		str    string
		expect int
	}{
		{
			str:    "tmmzuxt",
			expect: 5,
		},
		{
			str:    "ckilbkd",
			expect: 5,
		},
		{
			str:    "dvdf",
			expect: 3,
		},
		{
			str:    "aab",
			expect: 2,
		},
		{
			str:    "abcabcbb",
			expect: 3,
		},
		{
			str:    "bbbbb",
			expect: 1,
		},
		{
			str:    "pwwkew",
			expect: 3,
		},
		{
			str:    "",
			expect: 0,
		},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.str)
		if result != test.expect {
			t.Fatalf("failed result: %v, expect: %v", result, test.expect)
		}
	}
}
