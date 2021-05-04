package compare_the_triplets

func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			result[0] += 1
			continue
		} else if a[i] < b[i] {
			result[1] += 1
		}
	}

	return result
}
