package degree_of_an_array

type info struct {
	val      int
	count    int
	firstIdx int
	lastIdx  int
}

func findShortestSubArray(nums []int) int {
	infoMap := map[int]info{}

	for i, num := range nums {
		foundInfo, ok := infoMap[num]
		if !ok {
			info := info{
				val:      num,
				count:    1,
				firstIdx: i,
				lastIdx:  i,
			}
			infoMap[num] = info
		} else {
			foundInfo.count += 1
			foundInfo.lastIdx = i
			infoMap[num] = foundInfo
		}
	}

	max := 0
	for _, info := range infoMap {
		if info.count > max {
			max = info.count
		}
	}

	result := 50001
	for _, info := range infoMap {
		if info.count == max {
			subResult := info.lastIdx - info.firstIdx + 1
			if subResult < result {
				result = subResult
			}
		}
	}

	return result
}
