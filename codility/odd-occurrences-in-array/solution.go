package odd_occurrences_in_array

func Solution(A []int) int {
	paired := map[int]int{}

	for _, a := range A {
		paired[a] += 1
	}

	for k, v := range paired {
		if v%2 == 1 {
			return k
		}
	}

	return -1
}
