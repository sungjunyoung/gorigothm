package longest_palindrome

func longestPalindrome(s string) int {
	countMap := map[string]int{}

	result := 0
	for _, c := range s {
		countMap[string(c)] += 1
		result += 1
	}

	oddCount := 0
	for _, count := range countMap {
		if count%2 == 1 {
			oddCount += 1
		}
	}

	if oddCount < 2 {
		return result
	}

	return result - oddCount + 1
}
