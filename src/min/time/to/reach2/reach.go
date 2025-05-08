package reach2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minTimeToReach(moveTime [][]int) int {
	m := len(moveTime)
	n := len(moveTime[0])
	visited := make([][][2]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][2]int, n)
	}
	visited[0][0] = [2]int{0, 1}
	hp := h.New[[4]int]([][4]int{
		{0, 0, 0, 2},
	}, func(a, b [4]int) bool {
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	for hp.Len() != 0 {
		top := heap.Pop(&hp).([4]int)
		t, i, j, d := top[0], top[1], top[2], top[3]
		if t == visited[i][j][0] && 3-d == visited[i][j][1] {
			if i+1 == m && j+1 == n {
				return t
			}
			for _, dir := range dirs {
				ii, jj := dir[0], dir[1]
				ni, nj := i+ii, j+jj
				if !(ni == 0 && nj == 0) && ni >= 0 && ni < m && nj >= 0 && nj < n {
					nt := max(t, moveTime[ni][nj]) + 3 - d
					if visited[ni][nj][0] != 0 && (visited[ni][nj][0] < nt || (visited[ni][nj][0] == nt && visited[ni][nj][1] <= d)) {
						continue
					}
					visited[ni][nj] = [2]int{nt, d}
					heap.Push(&hp, [4]int{nt, ni, nj, 3 - d})
				}
			}
		}
	}
	return visited[m-1][n-1][0]
}
