package flips

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	e1 := m / 2
	e2 := n / 2
	var t1 int
	for j := 0; j < n; j++ {
		for i := 0; i < e1; i++ {
			if grid[i][j] != grid[m-1-i][j] {
				t1++
			}
		}
	}
	var t2 int
	for i := 0; i < m; i++ {
		for j := 0; j < e2; j++ {
			if grid[i][j] != grid[i][n-1-j] {
				t2++
			}
		}
	}
	if t1 < t2 {
		return t1
	}
	return t2
}
