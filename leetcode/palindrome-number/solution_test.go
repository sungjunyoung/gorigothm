package palindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		x      int
		expect bool
	}{
		{
			x:      121,
			expect: true,
		},
		{
			x:      -121,
			expect: false,
		},
		{
			x:      10,
			expect: false,
		},
		{
			x:      0,
			expect: true,
		},
	}

	for _, test := range tests {
		if isPalindrome(test.x) != test.expect {
			t.Fatal("failed")
		}
	}
}
