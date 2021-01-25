package reverse_integer

func reverse(x int) int {
	result := 0
	for {
		if x == 0 {
			break
		}

		result = (result * 10) + (x % 10)
		x = x / 10
	}

	if result < -2147483648 || result > 2147483648 {
		return 0
	}

	if x < 0 {
		return -result
	}

	return result
}
