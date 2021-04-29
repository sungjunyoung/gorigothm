package max_product_of_three

import (
	"sort"
)

func Solution(A []int) int {
	l := len(A)
	sort.Ints(A)

	max := A[l-1] * A[l-2] * A[l-3]
	if A[0]*A[1]*A[l-1] > max {
		max = A[0] * A[1] * A[l-1]
	}

	return max
}
