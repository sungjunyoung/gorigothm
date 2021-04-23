package frog_jmp

func Solution(X int, Y int, D int) int {
	result := (Y - X) / D
	others := (Y - X) % D
	if others != 0 {
		result += 1
	}

	return result
}
