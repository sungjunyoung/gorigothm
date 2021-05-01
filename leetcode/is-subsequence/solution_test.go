package is_subsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect bool
	}{
		{
			s:      "abc",
			t:      "ahbgdc",
			expect: true,
		},
		{
			s:      "axc",
			t:      "ahbgdc",
			expect: false,
		},
		{
			s:      "",
			t:      "ahbgdc",
			expect: true,
		},
	}

	for i, test := range tests {
		if isSubsequence(test.s, test.t) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
