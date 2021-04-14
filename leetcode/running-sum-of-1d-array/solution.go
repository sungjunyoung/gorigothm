package running_sum_of_1d_array

func runningSum(nums []int) []int {
	var result []int

	x := 0
	for _, num := range nums {
		x += num
		result = append(result, x)
	}

	return result
}
