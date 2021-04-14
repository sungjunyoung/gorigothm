package richest_customer_wealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	tests := []struct {
		accounts [][]int
		expect   int
	}{
		{
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			expect:   6,
		},
	}

	for _, test := range tests {
		if maximumWealth(test.accounts) != test.expect {
			t.Fatal("failed")
		}
	}
}
