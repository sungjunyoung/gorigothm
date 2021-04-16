package best_time_to_buy_and_sell_stock

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		expect int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}

	for _, test := range tests {
		if maxProfit(test.prices) != test.expect {
			t.Fatal("failed")
		}
	}
}
