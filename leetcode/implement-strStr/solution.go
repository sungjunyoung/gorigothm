package implement_strStr

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	needleLen := len(needle)

	for i := 0; i < len(haystack)-needleLen+1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
