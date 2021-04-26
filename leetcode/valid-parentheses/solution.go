package valid_parentheses

func isValid(s string) bool {
	opening := map[string]int{
		"(": 0,
		"{": 1,
		"[": 2,
	}
	closing := map[string]int{
		")": 0,
		"}": 1,
		"]": 2,
	}
	var stack []int

	for _, c := range s {
		if val, ok := opening[string(c)]; ok {
			stack = append(stack, val)
			continue
		}

		if val, ok := closing[string(c)]; ok {
			if len(stack) == 0 {
				return false
			}

			if val != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
