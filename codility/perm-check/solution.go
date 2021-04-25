package perm_check

func Solution(A []int) int {
	checker := map[int]bool{}

	for _, a := range A {
		checker[a] = true
	}

	for i := 1; i <= len(A); i++ {
		if !checker[i] {
			return 0
		}
	}

	return 1
}
