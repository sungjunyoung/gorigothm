package shortest_word_distance

import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	min := math.MaxInt64
	indexes := []int{math.MaxInt64, math.MaxInt64}

	for i, word := range wordsDict {
		if word == word1 {
			indexes[0] = i
			if indexes[1] != math.MaxInt64 {
				cur := abs(indexes[0] - indexes[1])
				if cur < min {
					min = cur
				}
			}
		}
		if word == word2 {
			indexes[1] = i
			if indexes[0] != math.MaxInt64 {
				cur := abs(indexes[0] - indexes[1])
				if cur < min {
					min = cur
				}
			}
		}
	}

	return min
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
