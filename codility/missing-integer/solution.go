package missing_integer

func Solution(A []int) int {
	max := A[0]
	helper := map[int]bool{}

	for _, a := range A {
		if a < 0 {
			continue
		}

		if a > max {
			max = a
		}
		helper[a] = true
	}

	if max < 0 {
		return 1
	}
	for i := 1; i <= max; i++ {
		if !helper[i] {
			return i
		}
	}

	return max + 1
}
