package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	var result []string

	for i := 1; i < n+1; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}

		result = append(result, strconv.Itoa(i))
	}

	return result
}
