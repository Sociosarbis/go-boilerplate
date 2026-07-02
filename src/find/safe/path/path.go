package path

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func findSafeWalk(grid [][]int, health int) bool {
	if grid[0][0] >= health {
		return false
	}
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return true
	}
	hp := h.New[[3]int]([][3]int{{0, 0, health - grid[0][0]}}, func(a, b [3]int) bool {
		return a[2] > b[2]
	})
	best := make([][]int, m)
	for i := 0; i < m; i++ {
		best[i] = make([]int, n)
	}
	for hp.Len() > 0 {
		top := heap.Pop(&hp).([3]int)
		if best[top[0]][top[1]] < top[2] {
			best[top[0]][top[1]] = top[2]
			for d := -1; d < 2; d += 2 {
				ni := top[0] + d
				if ni >= 0 && ni < m && top[2]-grid[ni][top[1]] > best[ni][top[1]] {
					if ni+1 == m && top[1]+1 == n {
						return true
					}
					heap.Push(&hp, [3]int{ni, top[1], top[2] - grid[ni][top[1]]})
				}
				nj := top[1] + d
				if nj >= 0 && nj < n && top[2]-grid[top[0]][nj] > best[top[0]][nj] {
					if top[0]+1 == m && nj+1 == n {
						return true
					}
					heap.Push(&hp, [3]int{top[0], nj, top[2] - grid[top[0]][nj]})
				}
			}
		}
	}
	return false
}
