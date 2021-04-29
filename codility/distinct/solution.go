package distinct

func Solution(A []int) int {
	counter := map[int]bool{}

	for _, a := range A {
		counter[a] = true
	}

	return len(counter)
}
