package defanging_an_ip_address

import "testing"

func Test_defangIPaddr(t *testing.T) {
	tests := []struct {
		address string
		expect  string
	}{
		{
			address: "1.1.1.1",
			expect:  "1[.]1[.]1[.]1",
		},
		{
			address: "255.100.50.0",
			expect:  "255[.]100[.]50[.]0",
		},
	}

	for _, test := range tests {
		if defangIPaddr(test.address) != test.expect {
			t.Fatal("failed")
		}
	}
}
