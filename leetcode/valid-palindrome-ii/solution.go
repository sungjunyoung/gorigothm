package valid_palindrome_ii

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			left, right := false, false
			left = isPalindrome(s[i+1 : j+1])
			right = isPalindrome(s[i:j])
			if !left && !right {
				return false
			}
		}
		i += 1
		j -= 1
	}

	return true
}

func isPalindrome(s string) bool {
	mid := len(s) / 2

	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
