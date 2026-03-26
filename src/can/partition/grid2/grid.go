package grid2

func canPartitionGrid(grid [][]int) bool {

	var sum int64
	var max int
	m, n := len(grid), len(grid[0])
	for _, row := range grid {
		for _, cell := range row {
			sum += int64(cell)
			if cell > max {
				max = cell
			}
		}
	}

	var temp int64
	var diff int64
	visited := make([]bool, max+1)
	if m == 1 {
		for i := 0; i < n; i++ {
			cell := grid[0][i]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff == 0 {
				return true
			}
			if diff >= 0 && diff <= int64(max) && (visited[diff] && int64(grid[0][0]) == diff || int64(grid[0][i]) == diff) {
				return true
			}
		}
		temp = 0
		visited = make([]bool, max+1)
		for i := n - 1; i >= 0; i-- {
			cell := grid[0][i]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff == 0 {
				return true
			}
			if diff >= 0 && diff <= int64(max) && (visited[diff] && int64(grid[0][n-1]) == diff || int64(grid[0][i]) == diff) {
				return true
			}
		}
		return false
	}

	if n == 1 {
		for i := 0; i < m; i++ {
			cell := grid[i][0]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff == 0 {
				return true
			}
			if diff >= 0 && diff <= int64(max) && (visited[diff] && int64(grid[0][0]) == diff || int64(grid[i][0]) == diff) {
				return true
			}
		}
		temp = 0
		visited = make([]bool, max+1)
		for i := m - 1; i >= 0; i-- {
			cell := grid[i][0]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff == 0 {
				return true
			}
			if diff >= 0 && diff <= int64(max) && (visited[diff] && int64(grid[m-1][0]) == diff || int64(grid[i][0]) == diff) {

				return true
			}
		}
		return false
	}

	for i, row := range grid {
		for _, cell := range row {
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff > int64(max) {
				break
			}
		}
		if diff == 0 {
			return true
		}
		if diff >= 0 && diff <= int64(max) && ((i != 0 && visited[diff]) || int64(grid[0][0]) == diff || int64(grid[0][n-1]) == diff) {
			return true
		}
		if diff > int64(max) {
			break
		}
	}
	visited = make([]bool, max+1)
	temp = 0
	for i := m - 1; i >= 0; i-- {
		row := grid[i]
		for _, cell := range row {
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff > int64(max) {
				break
			}
		}
		if diff == 0 {
			return true
		}
		if diff >= 0 && diff <= int64(max) && ((i != m-1 && visited[diff]) || int64(grid[m-1][0]) == diff || int64(grid[m-1][n-1]) == diff) {
			return true
		}
		if diff > int64(max) {
			break
		}
	}
	visited = make([]bool, max+1)
	temp = 0
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			cell := grid[i][j]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff > int64(max) {
				break
			}
		}
		if diff == 0 {
			return true
		}
		if diff >= 0 && diff <= int64(max) && ((j != 0 && visited[diff]) || int64(grid[0][0]) == diff || int64(grid[m-1][0]) == diff) {
			return true
		}
		if diff > int64(max) {
			break
		}
	}
	visited = make([]bool, max+1)
	temp = 0
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < m; i++ {
			cell := grid[i][j]
			temp += int64(cell)
			visited[cell] = true
			diff = temp*2 - sum
			if diff > int64(max) {
				break
			}
		}
		if diff == 0 {
			return true
		}
		if diff >= 0 && diff <= int64(max) && ((j != n-1 && visited[diff]) || int64(grid[0][n-1]) == diff || int64(grid[m-1][n-1]) == diff) {
			return true
		}
		if diff > int64(max) {
			break
		}
	}
	return false
}
