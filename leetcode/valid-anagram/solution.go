package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	anaMap := map[rune]int{}

	for _, r := range s {
		anaMap[r] += 1
	}

	for _, r := range t {
		if _, ok := anaMap[r]; ok {
			anaMap[r]--
			if anaMap[r] == 0 {
				delete(anaMap, r)
			}
		} else {
			return false
		}
	}

	return len(anaMap) == 0
}
