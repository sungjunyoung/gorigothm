package fish

func Solution(A []int, B []int) int {
	stack := []int{}
	upperAliveCount := 0
	for i := range A {
		if B[i] == 1 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			if A[i] > A[cur] {
				stack = stack[0 : len(stack)-1]
			} else {
				break
			}
		}

		if len(stack) == 0 {
			upperAliveCount++
		}
	}

	return len(stack) + upperAliveCount
}
