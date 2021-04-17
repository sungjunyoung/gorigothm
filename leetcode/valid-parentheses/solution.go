package valid_parentheses

func isValid(s string) bool {
	var stack []string

	for _, r := range s {
		switch string(r) {
		case "(":
			stack = append(stack, "(")
		case "{":
			stack = append(stack, "{")
		case "[":
			stack = append(stack, "[")
		}

		if len(stack) == 0 {
			return false
		}
		switch string(r) {
		case ")":
			if stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "}":
			if stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "]":
			if stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
