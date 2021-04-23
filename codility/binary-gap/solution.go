package binary_gap

func Solution(N int) int {
	var binaries []int
	for N != 0 {
		binaries = append(binaries, N%2)
		N = N / 2
	}

	result := 0
	counter := 0
	for i := len(binaries) - 1; i >= 0; i-- {
		if binaries[i] == 1 {
			if counter > result {
				result = counter
			}
			counter = 0
			continue
		}

		counter += 1
	}

	return result
}
