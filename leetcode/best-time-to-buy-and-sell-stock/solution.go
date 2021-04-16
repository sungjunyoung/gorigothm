package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	min := prices[0]
	result := 0
	for _, price := range prices {
		if price < min {
			min = price
		}

		if price-min > result {
			result = price - min
		}
	}
	return result
}
