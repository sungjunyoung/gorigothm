package two_sum

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, current := range nums {
		match := target - current
		if _, ok := hash[match]; ok {
			return []int{i, hash[match]}
		}

		hash[current] = i
	}
	return []int{}
}
