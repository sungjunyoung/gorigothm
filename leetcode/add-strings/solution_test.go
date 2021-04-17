package add_strings

import "testing"

func Test_addStrings(t *testing.T) {
	tests := []struct {
		num1   string
		num2   string
		expect string
	}{
		{
			num1:   "11",
			num2:   "123",
			expect: "134",
		},
		{
			num1:   "456",
			num2:   "77",
			expect: "533",
		},
		{
			num1:   "0",
			num2:   "0",
			expect: "0",
		},
		{
			num1:   "99",
			num2:   "99",
			expect: "198",
		},
	}

	for _, test := range tests {
		if addStrings(test.num1, test.num2) != test.expect {
			t.Fatal("failed")
		}
	}
}
