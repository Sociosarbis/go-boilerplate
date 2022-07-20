package grid

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	tempGrid := make([][]int, m)
	for i := range tempGrid {
		tempGrid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := i*n + j + k
			tempGrid[v/n%m][v%n] = grid[i][j]
		}
	}
	return tempGrid
}
