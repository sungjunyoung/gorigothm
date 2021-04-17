package decompress_run_length_encoded_list

func decompressRLElist(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}
