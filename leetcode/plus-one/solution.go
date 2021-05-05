package plus_one

func plusOne(digits []int) []int {
	l := len(digits)

	if digits[l-1] < 9 {
		digits[l-1] += 1
		return digits
	}

	digits[l-1] += 1
	additional := 0
	for i := l - 1; i >= 0; i-- {
		digit := digits[i] + additional
		if digit == 10 {
			digits[i] = 0
			additional = 1
		} else {
			digits[i] = digit
			additional = 0
		}
	}

	if additional == 1 {
		digits = append([]int{1}, digits...)
		return digits
	}

	return digits
}
