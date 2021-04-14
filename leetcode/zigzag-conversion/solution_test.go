package zigzag_conversion

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		expect  string
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			expect:  "PAHNAPLSIIGYIR",
		},
		{
			s:       "A",
			numRows: 1,
			expect:  "A",
		},
		{
			s:       "AB",
			numRows: 1,
			expect:  "AB",
		},
		{
			s:       "ABC",
			numRows: 2,
			expect:  "ACB",
		},
	}

	for _, test := range tests {
		if test.expect != convert(test.s, test.numRows) {
			t.Fatal("failed")
		}
	}
}
