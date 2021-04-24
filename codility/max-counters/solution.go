package max_counters

func Solution(N int, A []int) []int {
	result := make([]int, N)

	max := 0
	tmpMax := 0
	for _, a := range A {

		if a > N {
			max = tmpMax
			continue
		}

		if result[a-1] < max {
			result[a-1] = max
		}

		result[a-1] += 1
		if result[a-1] > tmpMax {
			tmpMax = result[a-1]
		}
	}

	for i := range result {
		if result[i] < max {
			result[i] = max
		}
	}

	return result
}
