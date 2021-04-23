package perm_missing_elem

func Solution(A []int) int {
	finder := map[int]bool{}
	for i := 1; i <= len(A)+1; i++ {
		finder[i] = false
	}

	for _, a := range A {
		finder[a] = true
	}

	for k, v := range finder {
		if !v {
			return k
		}
	}

	return -1
}
