package min_avg_two_slice

func Solution(A []int) int {
	min := float64(A[0]+A[1]) / 2
	length := len(A)
	if length == 2 {
		return 0
	}

	minIndex := 0
	for i := 2; i < length; i++ {
		avg := float64(A[i-2]+A[i-1]+A[i]) / 3
		if avg < min {
			min = avg
			minIndex = i - 2
		}

		avg = float64(A[i-1]+A[i]) / 2
		if avg < min {
			min = avg
			minIndex = i - 1
		}
	}
	return minIndex
}
