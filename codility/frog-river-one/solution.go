package frog_river_one

func Solution(X int, A []int) int {
	checker := map[int]bool{}
	for i, a := range A {
		checker[a] = true

		if len(checker) == X {
			return i
		}
	}
	return -1
}
