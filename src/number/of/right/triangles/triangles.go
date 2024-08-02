package triangles

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	n := len(grid[0])
	prefix_h := make([]int, m)
	prefix_v := make([]int, n)

	for i, row := range grid {
		for _, v := range row {
			prefix_h[i] += v
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			prefix_v[j] += grid[i][j]
		}
	}

	var ret int64

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				ret += int64(prefix_h[i]-1) * int64(prefix_v[j]-1)
			}
		}
	}

	return ret
}
