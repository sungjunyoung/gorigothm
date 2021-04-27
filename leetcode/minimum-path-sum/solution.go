package minimum_path_sum

import "math"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := range grid {
		for j := range grid[i] {
			up := math.MaxInt64
			left := math.MaxInt64

			if i-1 >= 0 {
				up = dp[i-1][j]
			}
			if j-1 >= 0 {
				left = dp[i][j-1]
			}

			min := up
			if left < min {
				min = left
			}

			if min == math.MaxInt64 {
				min = 0
			}

			dp[i][j] = grid[i][j] + min
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
