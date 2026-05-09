package grid

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func move(i, j, k, m, n int) (int, int) {
	k %= (m + n - 2) * 2
	var d int
	for k > 0 {
		// 向左
		if i == 0 && j != 0 {
			d = min(j, k)
			j -= d
			// 向下
		} else if j == 0 && i+1 != m {
			d = min(k, m-1-i)
			i += d
			// 向右
		} else if i+1 == m && j+1 != n {
			d = min(k, n-1-j)
			j += d
			// 向上
		} else {
			d = min(k, i)
			i -= d
		}
		k -= d
	}
	return i, j
}

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	out := make([][]int, m)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
	}
	l := min(m, n) / 2
	for d := 0; d < l; d++ {
		cm, cn := m-d*2, n-d*2
		for i := 0; i < cn; i++ {
			ni, nj := move(0, i, k, cm, cn)
			out[ni+d][nj+d] = grid[d][i+d]
			ni, nj = move(cm-1, i, k, cm, cn)
			out[ni+d][nj+d] = grid[cm-1+d][i+d]
		}
		for i := cm - 1; i > 0; i-- {
			ni, nj := move(i, 0, k, cm, cn)
			out[ni+d][nj+d] = grid[i+d][d]
			ni, nj = move(i, cn-1, k, cm, cn)
			out[ni+d][nj+d] = grid[i+d][cn-1+d]
		}
	}
	return out
}
