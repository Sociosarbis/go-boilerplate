package values

func abs(a int, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	suffix := make([][]int, m+1)
	for i := range suffix {
		suffix[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		pos := [2]int{m - 1, i}
		var mask int64
		var c int
		for pos[0] >= 0 && pos[1] >= 0 {
			var v int64 = 1 << grid[pos[0]][pos[1]]
			if mask&v == 0 {
				mask |= v
				c++
			}
			suffix[pos[0]][pos[1]] = c
			pos[0]--
			pos[1]--
		}
	}
	for i := m - 2; i >= 0; i-- {
		pos := [2]int{i, n - 1}
		var mask int64
		var c int
		for pos[0] >= 0 && pos[1] >= 0 {
			var v int64 = 1 << grid[pos[0]][pos[1]]
			if mask&v == 0 {
				mask |= v
				c++
			}
			suffix[pos[0]][pos[1]] = c
			pos[0]--
			pos[1]--
		}
	}
	for i := 0; i < n; i++ {
		var mask int64
		var c int
		pos := [2]int{0, i}
		for pos[0] < m && pos[1] < n {
			var v int64 = 1 << grid[pos[0]][pos[1]]
			grid[pos[0]][pos[1]] = abs(c, suffix[pos[0]+1][pos[1]+1])
			if mask&v == 0 {
				c++
				mask |= v
			}
			pos[0]++
			pos[1]++
		}
	}
	for i := 1; i < m; i++ {
		var mask int64
		var c int
		pos := [2]int{i, 0}
		for pos[0] < m && pos[1] < n {
			var v int64 = 1 << grid[pos[0]][pos[1]]
			grid[pos[0]][pos[1]] = abs(c, suffix[pos[0]+1][pos[1]+1])
			if mask&v == 0 {
				c++
				mask |= v
			}
			pos[0]++
			pos[1]++
		}
	}
	return grid
}
