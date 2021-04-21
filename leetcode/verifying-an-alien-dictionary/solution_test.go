package verifying_an_alien_dictionary

import "testing"

func Test_isAlienSorted(t *testing.T) {
	tests := []struct {
		words  []string
		order  string
		expect bool
	}{
		{
			words:  []string{"hello", "leetcode"},
			order:  "hlabcdefgijkmnopqrstuvwxyz",
			expect: true,
		},
		{
			words:  []string{"word", "world", "row"},
			order:  "worldabcefghijkmnpqstuvxyz",
			expect: false,
		},
	}

	for _, test := range tests {
		if isAlienSorted(test.words, test.order) != test.expect {
			t.Fatal("failed")
		}
	}
}
