package remove_vowels_from_a_string

func removeVowels(s string) string {
	filter := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	result := ""
	for _, r := range s {
		if !filter[r] {
			result += string(r)
		}
	}

	return result
}
