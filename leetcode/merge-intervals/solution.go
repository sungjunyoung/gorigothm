package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := [][]int{{intervals[0][0], intervals[0][1]}}
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		lastResult := results[len(results)-1]
		if lastResult[1] >= start {
			endMax := lastResult[1]
			if endMax < end {
				endMax = end
			}
			lastResult[1] = endMax
			results[len(results)-1] = lastResult
		} else {
			results = append(results, []int{start, end})
		}
	}

	return results
}
