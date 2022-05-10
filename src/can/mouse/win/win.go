package win

var DIRS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(grid *[]string, dp *[][][]int, cat, mouse, catJump, mouseJump, turn int) int {
	if turn >= len((*dp)[0][0]) {
		return 2
	}
	if (*dp)[cat][mouse][turn] != 0 {
		return (*dp)[cat][mouse][turn]
	}

	m := len(*grid)
	n := len((*grid)[0])
	if turn&1 == 0 {
		i := mouse / n
		j := mouse % n
		for k := 0; k < len(DIRS); k++ {
			for l := 0; l <= mouseJump; l++ {
				nextI := i + DIRS[k][0]*l
				nextJ := j + DIRS[k][1]*l
				if nextI < 0 || nextJ < 0 || nextI >= m || nextJ >= n || (*grid)[nextI][nextJ] == '#' {
					break
				}

				nextMouse := nextI*n + nextJ
				if (*grid)[nextI][nextJ] == 'F' || (nextMouse != cat && dfs(grid, dp, cat, nextMouse, catJump, mouseJump, turn+1) == 1) {
					(*dp)[cat][mouse][turn] = 1
					return 1
				}
			}
		}
		(*dp)[cat][mouse][turn] = 2
		return 2

	} else {
		i := cat / n
		j := cat % n
		for k := 0; k < len(DIRS); k++ {
			for l := 0; l <= catJump; l++ {
				nextI := i + DIRS[k][0]*l
				nextJ := j + DIRS[k][1]*l
				if nextI < 0 || nextJ < 0 || nextI >= m || nextJ >= n || (*grid)[nextI][nextJ] == '#' {
					break
				}
				nextCat := nextI*n + nextJ
				if (*grid)[nextI][nextJ] == 'F' || nextCat == mouse || dfs(grid, dp, nextCat, mouse, catJump, mouseJump, turn+1) == 2 {
					(*dp)[cat][mouse][turn] = 2
					return 2
				}
			}
		}
		(*dp)[cat][mouse][turn] = 1
		return 1
	}
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	count := len(grid) * len(grid[0])
	dp := make([][][]int, count)
	cat := 0
	mouse := 0
	maxTurn := count
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'C' {
				cat = i*len(grid[0]) + j
			} else if cell == 'M' {
				mouse = i*len(grid[0]) + j
			} else if cell == '#' {
				maxTurn--
			}
		}
	}
	for i := range dp {
		dp[i] = make([][]int, count)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxTurn*2)
		}
	}
	return dfs(&grid, &dp, cat, mouse, catJump, mouseJump, 0) == 1
}
