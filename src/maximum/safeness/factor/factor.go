package factor

func dfs(factors [][]int, m, n, i, j, min int) bool {
	if i+1 == m && j+1 == n {
		return true
	}
	factors[i][j] = -1 - factors[i][j]
	for d := -1; d < 2; d += 2 {
		ni := i + d
		nj := j + d
		if ni >= 0 && ni < m && factors[ni][j] >= min && dfs(factors, m, n, ni, j, min) {
			return true
		}
		if nj >= 0 && nj < n && factors[i][nj] >= min && dfs(factors, m, n, i, nj, min) {
			return true
		}
	}
	return false
}

func maximumSafenessFactor(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	factors := make([][]int, m)
	for i := 0; i < m; i++ {
		factors[i] = make([]int, n)
	}
	bfs := [][2]int{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				bfs = append(bfs, [2]int{i, j})
			}
		}
	}
	l := len(bfs)
	var s int
	for l > 0 {
		s++
		for k := 0; k < l; k++ {
			i, j := bfs[k][0], bfs[k][1]
			for d := -1; d < 2; d += 2 {
				ni := i + d
				nj := j + d
				if ni >= 0 && ni < m && grid[ni][j] == 0 && factors[ni][j] == 0 {
					factors[ni][j] = s
					bfs = append(bfs, [2]int{ni, j})
				}
				if nj >= 0 && nj < n && grid[i][nj] == 0 && factors[i][nj] == 0 {
					factors[i][nj] = s
					bfs = append(bfs, [2]int{i, nj})
				}
			}
		}
		bfs = bfs[l:]
		l = len(bfs)
	}
	l = 0
	r := factors[0][0]
	if r > factors[m-1][n-1] {
		r = factors[m-1][n-1]
	}

	var out int
	for l <= r {
		mid := (l + r) >> 1
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if factors[i][j] < 0 {
					factors[i][j] = -factors[i][j] - 1
				}
			}
		}
		if dfs(factors, m, n, 0, 0, mid) {
			out = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return out
}
