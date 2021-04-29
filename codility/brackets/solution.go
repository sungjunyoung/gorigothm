package brackets

func Solution(S string) int {
	openMap := map[string]int{
		"(": 0,
		"{": 1,
		"[": 2,
	}
	closeMap := map[string]int{
		")": 0,
		"}": 1,
		"]": 2,
	}

	var stack []int
	for _, s := range S {
		if val, ok := openMap[string(s)]; ok {
			stack = append(stack, val)
			continue
		}

		if val, ok := closeMap[string(s)]; ok {
			if len(stack) == 0 {
				return 0
			}

			if val == stack[len(stack)-1] {
				stack = stack[0 : len(stack)-1]
			} else {
				return 0
			}
		}
	}

	if len(stack) != 0 {
		return 0
	}

	return 1
}
