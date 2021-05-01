package majority_element

func majorityElement(nums []int) int {
	counter := map[int]int{} // num:count

	for _, num := range nums {
		counter[num]++
	}

	max := -1
	result := nums[0]
	for num, count := range counter {
		if count > max {
			max = count
			result = num
		}
	}

	return result
}
