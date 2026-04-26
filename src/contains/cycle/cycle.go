package cycle

func check(grid [][]byte, visited [][]bool, i, j, pi, pj int) bool {
	visited[i][j] = true
	for di := -1; di < 2; di++ {
		ni := i + di
		if ni < 0 || ni == len(grid) {
			continue
		}
		var s, e int
		if di == 0 {
			s, e = -1, 2
		} else {
			s, e = 0, 1
		}
		for dj := s; dj < e; dj += 2 {
			nj := j + dj
			if nj < 0 || nj == len(grid[0]) || (ni == pi && nj == pj) || grid[i][j] != grid[ni][nj] {
				continue
			}
			if visited[ni][nj] {
				return true
			}
			if check(grid, visited, ni, nj, i, j) {
				return true
			}
		}
	}
	return false
}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if check(grid, visited, i, j, -1, -1) {
					return true
				}
			}
		}
	}
	return false
}
