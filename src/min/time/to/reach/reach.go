package reach

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

const maxValue int = 1e9 + 2500

func minTimeToReach(moveTime [][]int) int {
	m := len(moveTime)
	n := len(moveTime[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	bfs := [][3]int{
		{0, 0, 0},
	}
	visited[0][0] = true
	var ct int
	c := len(bfs)
	var nextTime int = maxValue
	for c != 0 {
		for index := 0; index < c; index++ {
			t, i, j := bfs[index][0], bfs[index][1], bfs[index][2]
			if t > ct {
				if t < nextTime {
					nextTime = t
				}
				bfs = append(bfs, bfs[index])
			} else {
				for _, dir := range dirs {
					ii, jj := dir[0], dir[1]
					ni, nj := i+ii, j+jj
					if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] {
						nt := max(t, moveTime[ni][nj]) + 1
						if ni+1 == m && nj+1 == n {
							return nt
						}
						visited[ni][nj] = true
						bfs = append(bfs, [3]int{nt, ni, nj})
						if nt < nextTime {
							nextTime = nt
						}
					}
				}
			}
		}
		ct = nextTime
		nextTime = maxValue
		bfs = bfs[c:]
		c = len(bfs)
	}
	return ct
}
