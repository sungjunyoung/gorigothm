package happy_number

func isHappy(n int) bool {
	dupCheck := map[int]bool{}

	digitSquares := n
	for {
		digitSquares = calcDigitSquares(digitSquares)
		if dupCheck[digitSquares] {
			return false
		}
		dupCheck[digitSquares] = true

		if digitSquares == 1 {
			return true
		}
	}
}

func calcDigitSquares(n int) int {
	result := 0
	for n > 0 {
		tmp := n % 10
		n = n / 10
		result += tmp * tmp
	}

	return result
}
