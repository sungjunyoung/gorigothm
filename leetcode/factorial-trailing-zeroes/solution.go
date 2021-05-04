package factorial_trailing_zeroes

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	return (n / 5) + trailingZeroes(n/5)
}
