package island

func dfs(grid [][]int, i, j, id int) int {
	if i >= 0 && j >= 0 && i < len(grid) && j < len(grid) && grid[i][j] == 1 {
		grid[i][j] = id
		ret := 1
		return ret +
			dfs(grid, i-1, j, id) +
			dfs(grid, i+1, j, id) +
			dfs(grid, i, j-1, id) +
			dfs(grid, i, j+1, id)
	}
	return 0
}

func largestIsland(grid [][]int) int {
	groups := map[int]int{}
	id := 2
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				groups[id] = dfs(grid, i, j, id)
				id++
			}
		}
	}
	ret := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				neighborGroups := make(map[int]bool, 3)
				temp := 0
				if i > 0 {
					if grid[i-1][j] != 0 {
						neighborGroups[grid[i-1][j]] = true
						temp += groups[grid[i-1][j]]
					}
				}
				if i+1 < len(grid) {
					if grid[i+1][j] != 0 {
						if _, ok := neighborGroups[grid[i+1][j]]; !ok {
							neighborGroups[grid[i+1][j]] = true
							temp += groups[grid[i+1][j]]
						}
					}
				}
				if j > 0 {
					if grid[i][j-1] != 0 {
						if _, ok := neighborGroups[grid[i][j-1]]; !ok {
							neighborGroups[grid[i][j-1]] = true
							temp += groups[grid[i][j-1]]
						}
					}
				}
				if j+1 < len(grid) {
					if grid[i][j+1] != 0 {
						if _, ok := neighborGroups[grid[i][j+1]]; !ok {
							temp += groups[grid[i][j+1]]
						}
					}
				}
				if temp+1 > ret {
					ret = temp + 1
				}
			}
		}
	}
	for _, count := range groups {
		if count > ret {
			ret = count
		}
	}
	return ret
}
