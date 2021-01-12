package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0

	m := map[string]*int{}
	length := 0
	for i := range s {
		tmpI := i
		if m[string(s[i])] != nil {
			if length > res {
				res = length
			}

			length = i - *m[string(s[i])]

			for k, v := range m {
				if *v < *m[string(s[i])] {
					delete(m, k)
				}
			}
			m[string(s[i])] = &tmpI

			continue
		}

		length += 1
		m[string(s[i])] = &tmpI

		if i == len(s)-1 {
			if length > res {
				res = length
			}
		}
	}

	return res
}
