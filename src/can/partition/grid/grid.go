package grid

func canPartitionGrid(grid [][]int) bool {
	var sum int64
	m, n := len(grid), len(grid[0])
	for _, row := range grid {
		for _, cell := range row {
			sum += int64(cell)
		}
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2

	var temp int64
	for _, row := range grid {
		for j, cell := range row {
			temp += int64(cell)
			if temp == sum {
				if j == n-1 {
					return true
				}
			}
			if temp >= sum {
				break
			}
		}
	}
	temp = 0
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			temp += int64(grid[i][j])
			if temp == sum {
				if i == m-1 {
					return true
				}
			}
			if temp >= sum {
				break
			}
		}
	}
	return false
}
