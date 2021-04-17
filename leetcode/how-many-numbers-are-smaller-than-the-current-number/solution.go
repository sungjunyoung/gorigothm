package how_many_numbers_are_smaller_than_the_current_number

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	cpy := make([]int, len(nums))
	lookup := map[int]int{}

	copy(cpy, nums)
	sort.Ints(cpy)
	for i, c := range cpy {
		_, ok := lookup[c]
		if !ok {
			lookup[c] = i
		}
	}

	var result []int
	for _, v := range nums {
		result = append(result, lookup[v])
	}

	return result
}
