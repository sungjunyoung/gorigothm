package longest_palindromic_substring

func longestPalindrome(s string) string {
	result := string(s[0])

	for i := range s {
		if i == 0 {
			continue
		}

		substr := findFrontPalindrome(s, i)
		if len(substr) > len(result) {
			result = substr
		}

		substr = findSidePalindrome(s, i)
		if len(substr) > len(result) {
			result = substr
		}
	}
	return result
}

func findFrontPalindrome(s string, pivot int) string {
	start := pivot - 1
	end := pivot
	for {
		if start < 0 || end > len(s)-1 || s[start] != s[end] {
			return s[start+1 : end-1+1]
		}

		start -= 1
		end += 1
	}
}

func findSidePalindrome(s string, pivot int) string {
	start := pivot - 1
	end := pivot + 1
	for {
		if start < 0 || end > len(s)-1 || s[start] != s[end] {
			return s[start+1 : end-1+1]
		}

		start -= 1
		end += 1
	}
}
