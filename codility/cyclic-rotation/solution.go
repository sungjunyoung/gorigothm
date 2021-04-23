package cyclic_rotation

func Solution(A []int, K int) []int {
	result := make([]int, len(A))

	for i, a := range A {
		ri := i + K
		for ri >= len(result) {
			ri -= len(result)
		}

		result[ri] = a
	}

	return result
}
