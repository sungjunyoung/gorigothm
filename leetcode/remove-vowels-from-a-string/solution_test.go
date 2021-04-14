package remove_vowels_from_a_string

import "testing"

func Test_removeVowels(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{
			s:      "leetcodeisacommunityforcoders",
			expect: "ltcdscmmntyfrcdrs",
		},
		{
			s:      "aeiou",
			expect: "",
		},
	}

	for _, test := range tests {
		if removeVowels(test.s) != test.expect {
			t.Fatal("failed")
		}
	}
}
