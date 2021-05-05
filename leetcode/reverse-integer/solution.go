package reverse_integer

import "math"

func reverse(x int) int {
	tmp := x
	i := 1
	for tmp != 0 {
		i *= 10
		tmp /= 10
	}
	i /= 10

	result := 0
	for {
		if x == 0 {
			break
		}

		rest := x % 10
		x = x / 10

		result += rest * i
		i /= 10
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}
