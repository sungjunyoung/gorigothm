package kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	for _, candy := range candies {
		if maxCandy < candy {
			maxCandy = candy
		}
	}

	var result []bool
	for _, candy := range candies {
		result = append(result, maxCandy <= candy+extraCandies)
	}

	return result
}
