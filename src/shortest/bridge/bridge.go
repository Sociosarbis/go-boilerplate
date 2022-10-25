package bridge

func isValid(x, y, n int) bool {
	return x >= 0 && y >= 0 && x < n && y < n
}

func dfs(grid [][]int, i, j int, border *[][]int) {
	if isValid(i, j, len(grid)) {
		if grid[i][j] == 1 {
			grid[i][j] = 2
			add := false
			if i+1 < len(grid) {
				dfs(grid, i+1, j, border)
				if grid[i+1][j] == 0 {
					add = true
				}
			}
			if i-1 >= 0 {
				dfs(grid, i-1, j, border)
				if !add && grid[i-1][j] == 0 {
					add = true
				}
			}
			if j+1 < len(grid) {
				dfs(grid, i, j+1, border)
				if !add && grid[i][j+1] == 0 {
					add = true
				}
			}
			if j-1 >= 0 {
				dfs(grid, i, j-1, border)
				if !add && grid[i][j-1] == 0 {
					add = true
				}
			}
			if add {
				*border = append(*border, []int{i, j})
			}
		}
	}
}

func shortestBridge(grid [][]int) int {
	border := [][]int{}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				dfs(grid, i, j, &border)
				ret := 0
				for len(border) != 0 {
					n := len(border)
					for i := 0; i < n; i++ {
						x := border[i][0]
						y := border[i][1]
						for _, dir := range dirs {
							nextX, nextY := x+dir[0], y+dir[1]
							if isValid(nextX, nextY, len(grid)) {
								if grid[nextX][nextY] == 1 {
									return ret
								} else if grid[nextX][nextY] == 0 {
									grid[nextX][nextY] = 2
									border = append(border, []int{nextX, nextY})
								}
							}
						}
					}
					ret++
				}
				break
			}
		}
	}
	return 0
}
