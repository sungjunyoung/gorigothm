package move_zeros

func moveZeroes(nums []int) {
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			j := i
			for nums[j] == 0 {
				if j == len(nums)-1 {
					break
				}
				j++
			}

			for j != i {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				j--
			}
			i++
		} else {
			i++
		}
	}

	return
}
