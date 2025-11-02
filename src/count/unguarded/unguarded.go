package unguarded

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	visited := make([][]byte, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]byte, n)
	}
	for _, guard := range guards {
		r, c := guard[0], guard[1]
		visited[r][c] = 2
	}
	for _, wall := range walls {
		r, c := wall[0], wall[1]
		visited[r][c] = 3
	}
	for _, guard := range guards {
		r, c := guard[0], guard[1]
		for i := r - 1; i >= 0; i-- {
			if visited[i][c] == 0 {
				visited[i][c] = 1
			} else if visited[i][c] != 1 {
				break
			}
		}
		for i := r + 1; i < m; i++ {
			if visited[i][c] == 0 {
				visited[i][c] = 1
			} else if visited[i][c] != 1 {
				break
			}
		}
		for i := c - 1; i >= 0; i-- {
			if visited[r][i] == 0 {
				visited[r][i] = 1
			} else if visited[r][i] != 1 {
				break
			}
		}
		for i := c + 1; i < n; i++ {
			if visited[r][i] == 0 {
				visited[r][i] = 1
			} else if visited[r][i] != 1 {
				break
			}
		}
	}
	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == 0 {
				count++
			}
		}
	}
	return count
}
