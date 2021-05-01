package first_unique_character_in_a_string

import "testing"

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{
			s:      "leetcode",
			expect: 0,
		},
	}

	for _, test := range tests {
		if firstUniqChar(test.s) != test.expect {
			t.Error("failed")
		}
	}
}
