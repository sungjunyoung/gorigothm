package count_and_say

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		n      int
		expect string
	}{
		{
			n:      1,
			expect: "1",
		},
		{
			n:      4,
			expect: "1211",
		},
		{
			n:      5,
			expect: "111221",
		},
	}

	for _, test := range tests {
		assert.Equal(t, countAndSay(test.n), test.expect)
	}
}
