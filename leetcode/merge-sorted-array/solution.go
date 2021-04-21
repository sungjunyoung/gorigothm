package merge_sorted_array

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	for j < n {
		if nums1[i] == 0 {
			nums1[i] = nums2[j]
			j++
		}
		i++
	}
	sort.Ints(nums1)
}
