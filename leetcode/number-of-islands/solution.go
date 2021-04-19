package number_of_islands

type pair struct {
	i int
	j int
}

func numIslands(grid [][]byte) int {
	visits := make([][]bool, len(grid))
	for i := range visits {
		visits[i] = make([]bool, len(grid[i]))
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visits[i][j] {
				continue
			}

			if grid[i][j] == '1' {
				var stack []pair
				stack = append(stack, pair{i: i, j: j})
				visits[i][j] = true
				for len(stack) != 0 {
					current := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					visits[current.i][current.j] = true

					if current.i-1 >= 0 && grid[current.i-1][current.j] == '1' && !visits[current.i-1][current.j] {
						stack = append(stack, pair{i: current.i - 1, j: current.j})
					}
					if current.j-1 >= 0 && grid[current.i][current.j-1] == '1' && !visits[current.i][current.j-1] {
						stack = append(stack, pair{i: current.i, j: current.j - 1})
					}
					if current.i+1 < len(grid) && grid[current.i+1][current.j] == '1' && !visits[current.i+1][current.j] {
						stack = append(stack, pair{i: current.i + 1, j: current.j})
					}
					if current.j+1 < len(grid[i]) && grid[current.i][current.j+1] == '1' && !visits[current.i][current.j+1] {
						stack = append(stack, pair{i: current.i, j: current.j + 1})
					}
				}
				result += 1
			}
		}
	}

	return result
}
