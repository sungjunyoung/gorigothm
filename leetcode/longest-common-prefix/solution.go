package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	return lcp(strs, 0, len(strs)-1)
}

func lcp(strs []string, l int, r int) string {
	if l == r {
		return strs[l]
	}

	mid := (l + r) / 2
	lResult := lcp(strs, l, mid)
	rResult := lcp(strs, mid+1, r)
	return cp(lResult, rResult)
}

func cp(left string, right string) string {
	minLen := len(left)
	if len(right) < minLen {
		minLen = len(right)
	}

	for i := 0; i < minLen; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}

	return left[:minLen]
}
