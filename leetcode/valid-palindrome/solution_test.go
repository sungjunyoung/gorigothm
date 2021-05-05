package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect bool
	}{
		{
			s:      "Marge, let's \"[went].\" I await {news} telegram.",
			expect: true,
		},
		{
			s:      "0p",
			expect: false,
		},
		{
			s:      "A man, a plan, a canal: Panama",
			expect: true,
		},
		{
			s:      "race a car",
			expect: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, isPalindrome(test.s))
	}
}
