package ransom_note

import "testing"

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		ransomNode string
		magazine   string
		expect     bool
	}{
		{
			ransomNode: "a",
			magazine:   "b",
			expect:     false,
		},
		{
			ransomNode: "aa",
			magazine:   "ab",
			expect:     false,
		},
		{
			ransomNode: "aa",
			magazine:   "aab",
			expect:     true,
		},
	}

	for i, test := range tests {
		if canConstruct(test.ransomNode, test.magazine) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
