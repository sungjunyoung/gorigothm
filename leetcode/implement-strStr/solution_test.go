package implement_strStr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expect   int
	}{
		{
			haystack: "aaaaa",
			needle:   "bba",
			expect:   -1,
		},
		{
			haystack: "hello",
			needle:   "ll",
			expect:   2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, strStr(test.haystack, test.needle), test.expect)
	}
}
