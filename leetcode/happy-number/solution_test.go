package happy_number

import "testing"

func Test_isHappy(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{
			n:      19,
			expect: true,
		},
		{
			n:      2,
			expect: false,
		},
	}

	for _, test := range tests {
		if isHappy(test.n) != test.expect {
			t.Fatal("failed")
		}
	}
}
