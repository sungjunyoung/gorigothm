package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	counter := map[rune]int{}

	for _, r := range ransomNote {
		counter[r]++
	}

	for _, r := range magazine {
		counter[r]--
	}

	for _, count := range counter {
		if count > 0 {
			return false
		}
	}

	return true
}
