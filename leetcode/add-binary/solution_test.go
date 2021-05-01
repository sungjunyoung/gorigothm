package add_binary

import "testing"

func Test_addBinary(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect string
	}{
		{
			a:      "11",
			b:      "1",
			expect: "100",
		},
		{
			a:      "1010",
			b:      "1011",
			expect: "10101",
		},
	}

	for i, test := range tests {
		if addBinary(test.a, test.b) != test.expect {
			t.Errorf("failed for test %d", i)
		}
	}
}
