package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	allProduct := 1
	var zeroIndexs []int
	for i, num := range nums {
		if num != 0 {
			allProduct *= num
		} else {
			zeroIndexs = append(zeroIndexs, i)
		}
	}

	results := make([]int, len(nums))
	if len(zeroIndexs) > 1 {
		return results
	}

	if len(zeroIndexs) == 1 {
		for _, zeroIndex := range zeroIndexs {
			results[zeroIndex] = allProduct
		}
		return results
	}

	for i, num := range nums {
		results[i] = allProduct / num
	}
	return results
}
