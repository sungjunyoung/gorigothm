package count_and_say

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	prevRes := countAndSay(n - 1)

	result := ""
	var cur = rune(prevRes[0])
	count := 1
	for i := 1; i < len(prevRes); i++ {
		if rune(prevRes[i]) != cur {
			result += fmt.Sprintf("%d%c", count, cur)
			cur = rune(prevRes[i])
			count = 1
		} else {
			count += 1
		}
	}

	result += fmt.Sprintf("%d%c", count, cur)

	return result
}
