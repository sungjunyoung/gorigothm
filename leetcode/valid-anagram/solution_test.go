package valid_anagram

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect bool
	}{
		{
			s:      "anagram",
			t:      "nagaram",
			expect: true,
		},
		{
			s:      "네이버",
			t:      "이네버",
			expect: true,
		},
	}

	for _, test := range tests {
		if isAnagram(test.t, test.s) != test.expect {
			t.Error("failed")
		}
	}
}
