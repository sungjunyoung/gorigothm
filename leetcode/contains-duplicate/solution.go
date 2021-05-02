package contains_duplicate

func containsDuplicate(nums []int) bool {
	dup := map[int]int{}
	for _, num := range nums {
		dup[num] += 1

		if dup[num] > 1 {
			return true
		}
	}

	return false
}
