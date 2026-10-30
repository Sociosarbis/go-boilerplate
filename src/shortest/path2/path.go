package path2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func shortestPath(n int, edges [][]int, labels string, k int) int {
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, k+1)
	}
	g := make([][][2]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], [2]int{edge[1], edge[2]})
	}
	hp := h.New([][3]int{[3]int{0, 0, 1}}, func(a, b [3]int) bool {
		return a[1] < b[1]
	})
	for hp.Len() != 0 {
		top := heap.Pop(&hp).([3]int)
		u, t, c := top[0], top[1], top[2]
		if t > visited[u][c] {
			continue
		}
		if u == n-1 {
			return t
		}
		var nc int
		for _, next := range g[u] {
			v, w := next[0], next[1]
			if labels[v] == labels[u] {
				if c+1 > k {
					continue
				}
				nc = c + 1
			} else {
				nc = 1
			}
			if visited[v][nc] == 0 || t+w < visited[v][nc] {
				visited[v][nc] = t + w
				heap.Push(&hp, [3]int{v, t + w, nc})
			}
		}
	}
	return -1
}
