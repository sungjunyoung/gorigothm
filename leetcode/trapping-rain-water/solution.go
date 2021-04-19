package trapping_rain_water

func trap(heights []int) int {
	if len(heights) < 2 {
		return 0
	}

	leftMaxs := make([]int, len(heights))
	for i := 1; i < len(leftMaxs); i++ {
		if leftMaxs[i-1] > heights[i-1] {
			leftMaxs[i] = leftMaxs[i-1]
		} else {
			leftMaxs[i] = heights[i-1]
		}
	}

	rightMaxs := make([]int, len(heights))
	for i := len(rightMaxs) - 2; i > 0; i-- {
		if rightMaxs[i+1] > heights[i+1] {
			rightMaxs[i] = rightMaxs[i+1]
		} else {
			rightMaxs[i] = heights[i+1]
		}
	}

	result := 0
	for i, height := range heights {
		lrMin := leftMaxs[i]
		if rightMaxs[i] < lrMin {
			lrMin = rightMaxs[i]
		}

		subResult := lrMin - height
		if subResult > 0 {
			result += subResult
		}
	}

	return result
}
