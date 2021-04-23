package tape_equilibrium

func Solution(A []int) int {
	lMax := make([]int, len(A))
	rMax := make([]int, len(A))

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		lMax[i] = sum
	}

	sum = 0
	for i := len(A) - 1; i >= 0; i-- {
		sum += A[i]
		rMax[i] = sum
	}

	result := abs(lMax[0] - rMax[1])
	for i := 1; i < len(A)-1; i++ {
		tmp := abs(lMax[i] - rMax[i+1])

		if tmp < result {
			result = tmp
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
