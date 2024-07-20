package moves3

func minimumMoves(grid [][]int) int {
	options := [][2]int{}
	destinations := [][2]int{}
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				destinations = append(destinations, [2]int{i, j})
			} else if v > 1 {
				options = append(options, [2]int{i, j})
			}
		}
	}
	return dfs(grid, destinations, options, 0, 0)
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func dfs(grid [][]int, destinations, options [][2]int, i, temp int) int {
	if i >= len(destinations) {
		return temp
	}
	ret := 32
	x1, y1 := destinations[i][0], destinations[i][1]
	for _, option := range options {
		x2, y2 := option[0], option[1]
		if grid[x2][y2] > 1 {
			grid[x2][y2]--
			res := dfs(grid, destinations, options, i+1, temp+abs(x1-x2)+abs(y1-y2))
			if res < ret {
				ret = res
			}
			grid[x2][y2]++
		}
	}
	return ret
}
