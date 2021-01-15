package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMedian(nums2)
	}
	if len(nums2) == 0 {
		return getMedian(nums1)
	}

	x := 0
	y := 0
	var arr []int

	for {
		if x == len(nums1) && y == len(nums2) {
			break
		}

		if x == len(nums1) {
			arr = append(arr, nums2[y])
			y += 1
			continue
		}

		if y == len(nums2) {
			arr = append(arr, nums1[x])
			x += 1
			continue
		}

		if nums1[x] > nums2[y] {
			arr = append(arr, nums2[y])
			y += 1
		} else {
			arr = append(arr, nums1[x])
			x += 1
		}
	}

	return getMedian(arr)
}

func getMedian(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}

	if len(arr)%2 != 0 {
		return float64(arr[len(arr)/2])
	}

	return float64(arr[len(arr)/2-1]+arr[len(arr)/2]) / 2
}
