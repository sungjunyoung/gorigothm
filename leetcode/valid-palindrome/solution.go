package valid_palindrome

func isPalindrome(s string) bool {
	val := []rune{}
	for _, r := range s {
		if r-'a' >= 0 && r-'a' < 26 {
			val = append(val, r)
		}
		if r-'A' >= 0 && r-'A' < 26 {
			val = append(val, r+32)
		}
		if r-'0' >= 0 && r-'0' < 10 {
			val = append(val, r)
		}
	}

	for i := 0; i < len(val)/2; i++ {
		if val[i] != val[len(val)-1-i] {
			return false
		}
	}

	return true
}
