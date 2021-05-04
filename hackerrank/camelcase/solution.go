package camelcase

func camelcase(s string) int32 {
	var result int32
	for i := range s {
		c := s[i] - 'a'
		if c >= 26 {
			result += 1
		}
	}

	return result + 1
}
