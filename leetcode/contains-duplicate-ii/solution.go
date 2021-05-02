package contains_duplicate_ii

type dupInfo struct {
	count int
	index int
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dups := map[int]*dupInfo{}

	for i, num := range nums {
		if d, ok := dups[num]; !ok {
			dups[num] = &dupInfo{
				count: 1,
				index: i,
			}
		} else {
			if i-d.index <= k {
				return true
			}

			dups[num].index = i
			dups[num].count += 1
		}
	}

	return false
}
