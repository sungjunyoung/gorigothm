package number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	result := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result += 1
			}
		}
	}

	return result
}
