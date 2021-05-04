package intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	counterMap := map[int]int{}
	for _, num := range nums1 {
		counterMap[num] += 1
	}

	result := []int{}
	for _, num := range nums2 {
		if counterMap[num] > 0 {
			result = append(result, num)
			counterMap[num] -= 1
		}
	}

	return result
}
