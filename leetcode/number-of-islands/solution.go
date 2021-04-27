package number_of_islands

func numIslands(grid [][]byte) int {
	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				traverse(i, j, grid)
				result += 1
			}
		}
	}

	return result
}

func traverse(i int, j int, grid [][]byte) {
	grid[i][j] = 'x'

	// up
	if i-1 >= 0 && grid[i-1][j] == '1' {
		traverse(i-1, j, grid)
	}
	// down
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		traverse(i+1, j, grid)
	}
	// left
	if j-1 >= 0 && grid[i][j-1] == '1' {
		traverse(i, j-1, grid)
	}
	// right
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		traverse(i, j+1, grid)
	}
}
