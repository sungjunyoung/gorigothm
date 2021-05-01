package is_subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for _, c := range t {
		if len(s) == 0 {
			return true
		}

		if int32(s[0]) == c {
			// dequeue
			if len(s) == 1 {
				s = ""
			} else {
				s = s[1:]
			}
		}
	}

	return len(s) == 0
}
