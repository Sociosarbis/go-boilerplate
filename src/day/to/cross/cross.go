package cross

var dirs [][2]int = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func isValid(cells [][]int, k, row, col int) bool {
	m, visited := make([][]bool, row), make([][]bool, row)
	for i := 0; i < row; i++ {
		m[i], visited[i] = make([]bool, col), make([]bool, col)
	}
	for i := 0; i < k; i++ {
		m[cells[i][0]-1][cells[i][1]-1] = true
	}
	bfs := make([][2]int, 0, col)
	for i := 0; i < col; i++ {
		if !m[0][i] {
			bfs = append(bfs, [2]int{0, i})
			visited[0][i] = true
		}
	}
	n := len(bfs)
	for len(bfs) != 0 {
		for i := 0; i < n; i++ {
			r, c := bfs[i][0], bfs[i][1]
			for _, dir := range dirs {
				rr, cc := r+dir[0], c+dir[1]
				if rr >= 0 && rr < row && cc >= 0 && cc < col {
					if !visited[rr][cc] && !m[rr][cc] {
						if rr+1 == row {
							return true
						}
						visited[rr][cc] = true
						bfs = append(bfs, [2]int{rr, cc})
					}
				}
			}
		}
		bfs = bfs[n:]
		n = len(bfs)
	}
	return false
}

func latestDayToCross(row int, col int, cells [][]int) int {
	var l int
	r := len(cells)
	var out int
	for l <= r {
		k := (l + r) >> 1
		if isValid(cells, k, row, col) {
			if k > out {
				out = k
			}
			l = k + 1
		} else {
			r = k - 1
		}
	}
	return out
}
