package single_number

func singleNumber(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}
