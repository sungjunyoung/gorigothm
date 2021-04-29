package nesting

func Solution(S string) int {
	stack := []string{}

	for _, s := range S {
		if string(s) == "(" {
			stack = append(stack, "(")
		} else {
			if len(stack) == 0 {
				return 0
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return 0
	}

	return 1
}
