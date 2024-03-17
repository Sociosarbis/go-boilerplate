package moves

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	prevCol := make([]bool, m)
	for i := 0; i < m; i++ {
		prevCol[i] = true
	}
loop:
	for i := 1; i < n; i++ {
		col := make([]bool, m)
		for j := 0; j < m; j++ {
			if prevCol[j] {
				if j > 0 {
					if grid[j][i-1] < grid[j-1][i] {
						col[j-1] = true
					}
				}
				if grid[j][i-1] < grid[j][i] {
					col[j] = true
				}
				if j+1 < m {
					if grid[j][i-1] < grid[j+1][i] {
						col[j+1] = true
					}
				}

			}
		}
		prevCol = col
		for i := 0; i < m; i++ {
			if prevCol[i] {
				continue loop
			}
		}
		return i - 1
	}
	return n - 1
}
