package matrix

const mask int = 12345

func constructProductMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	prefix := make([]int, m*n+1)
	suffix := make([]int, m*n+1)
	prefix[0], suffix[m*n] = 1, 1
	for i, row := range grid {
		for j, cell := range row {
			index := i*n + j
			prefix[index+1] = (prefix[index] * (cell % mask)) % mask
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			index := i*n + j
			suffix[index] = (suffix[index+1] * (grid[i][j] % mask)) % mask
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			grid[i][j] = (prefix[index] * suffix[index+1]) % mask
		}
	}
	return grid
}
