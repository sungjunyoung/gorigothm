package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	result := 0
	for _, banks := range accounts {
		sum := 0
		for _, bank := range banks {
			sum += bank
		}

		if result < sum {
			result = sum
		}
	}

	return result
}
