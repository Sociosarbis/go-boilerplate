package cost6

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minCost(n int, edges [][]int) int {
	g := make([][][2]int, n)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, 2 * w})
	}
	costs := make([]int, n)
	hp := h.New([][2]int{{0, 0}}, func(a, b [2]int) bool {
		return a[1] < b[1]
	})
	for hp.Len() > 0 {
		top := heap.Pop(&hp).([2]int)
		u, cost := top[0], top[1]
		if cost == costs[u] {
			for _, neighbor := range g[u] {
				v, w := neighbor[0], neighbor[1]
				if v != 0 && (costs[v] == 0 || cost+w < costs[v]) {
					costs[v] = cost + w
					heap.Push(&hp, [2]int{v, cost + w})
				}
			}
		}
	}
	if costs[n-1] == 0 {
		return -1
	}
	return costs[n-1]
}
