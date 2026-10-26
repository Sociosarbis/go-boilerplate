package islands

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	out := grid[i][j]
	grid[i][j] = 0
	for d := -1; d < 2; d += 2 {
		if i+d >= 0 && i+d < len(grid) {
			out += dfs(grid, i+d, j)
		}
		if j+d >= 0 && j+d < len(grid[0]) {
			out += dfs(grid, i, j+d)
		}
	}
	return out
}

func countIslands(grid [][]int, k int) int {
	var out int
	for i, row := range grid {
		for j, value := range row {
			if value != 0 {
				res := dfs(grid, i, j)
				if res%k == 0 {
					out++
				}
			}
		}
	}
	return out
}
