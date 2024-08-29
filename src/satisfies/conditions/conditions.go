package conditions

func satisfiesConditions(grid [][]int) bool {
	prev := 10
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < n; i++ {
		v := grid[0][i]
		if v == prev {
			return false
		}
		for j := 1; j < m; j++ {
			if grid[j][i] != v {
				return false
			}
		}
		prev = v
	}
	return true
}
