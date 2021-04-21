package verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	orderMap := map[string]int{"-1": 0}
	for i, o := range order {
		orderMap[string(o)] = i + 1
	}

	for i := 0; i < len(words)-1; i++ {
		wordA := words[i]
		wordB := words[i+1]

		maxLen := len(wordA)
		if len(wordB) > maxLen {
			maxLen = len(wordB)
		}
		for j := 0; j < maxLen; j++ {
			charA, charB := "-1", "-1"
			if len(wordA) > j {
				charA = string(wordA[j])
			}
			if len(wordB) > j {
				charB = string(wordB[j])
			}

			if charA == charB {
				continue
			}

			if orderMap[charA] > orderMap[charB] {
				return false
			}

			break
		}
	}

	return true
}
