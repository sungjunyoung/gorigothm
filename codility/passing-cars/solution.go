package passing_cars

func Solution(A []int) int {
	rightOneCountSums := make([]int, len(A))
	var zeroIndexs []int
	curOneCount := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 1 {
			curOneCount += 1
			rightOneCountSums[i] = curOneCount
			continue
		}

		rightOneCountSums[i] = curOneCount
		zeroIndexs = append(zeroIndexs, i)
	}

	result := 0
	for _, zeroIndex := range zeroIndexs {
		result += rightOneCountSums[zeroIndex]
		if result > 1000000000 {
			return -1
		}
	}

	return result
}
