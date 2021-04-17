package sign_of_the_product_of_an_array

func arraySign(nums []int) int {
	productResult := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num > 0 {
			productResult *= 1
		} else {
			productResult *= -1
		}

	}

	return productResult
}
