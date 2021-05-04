package remove_duplicates_from_sorted_array

import (
	"math"
	"sort"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur == nums[i] {
			nums[i] = math.MaxInt64
		} else {
			cur = nums[i]
		}
	}

	sort.Ints(nums)
	result := 0
	for i := range nums {
		if nums[i] == math.MaxInt64 {
			break
		}
		result++
	}
	return result
}
