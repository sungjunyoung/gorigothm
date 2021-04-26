package generate_parentheses

func generateParenthesis(n int) []string {
	var result []string
	traverse(n, 0, 0, "", &result)
	return result
}

func traverse(n int, left int, right int, str string, result *[]string) {
	if n == left && n == right {
		*result = append(*result, str)
		return
	}

	if left < n {
		traverse(n, left+1, right, str+"(", result)
	}
	if left > right {
		traverse(n, left, right+1, str+")", result)
	}
}
