package hamming_distance

func hammingDistance(x int, y int) int {
	xor := x ^ y
	result := 0
	for xor != 0 {
		result += xor & 1
		xor = xor >> 1
	}

	return result
}
