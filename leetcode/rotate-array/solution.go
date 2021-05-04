package rotate_array

func rotate(nums []int, k int) {
	orig := make([]int, len(nums))
	copy(orig, nums)
	k = k % len(nums)
	if k == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {
		targetIdx := i + k
		if targetIdx >= len(nums) {
			targetIdx -= len(nums)
		}

		nums[targetIdx] = orig[i]
	}
}
