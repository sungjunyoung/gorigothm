package maximum_subarray

func maxSubArray(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
		if nums[i] > sums[i] {
			sums[i] = nums[i]
		}

		if sums[i] > max {
			max = sums[i]
		}
	}

	return max
}
