package score2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type item struct {
	cost int64
	u    int
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	g := make([][][2]int, n)
	for i := 0; i < n; i++ {
		g[i] = [][2]int{}
	}
	var r int
	for _, edge := range edges {
		u, v, cost := edge[0], edge[1], edge[2]
		if !online[u] && online[v] {
			continue
		}
		g[u] = append(g[u], [2]int{v, cost})
		if cost > r {
			r = cost
		}
	}
	var l int
	out := -1
	initialCosts := make([]int64, n)
	for i := n - 1; i > 0; i-- {
		initialCosts[i] = -1
	}
	costs := make([]int64, n)
block:
	for l <= r {
		mid := (l + r) >> 1
		copy(costs, initialCosts)
		hp := h.New[item]([]item{{0, 0}}, func(a, b item) bool {
			return a.cost < b.cost
		})
		for hp.Len() > 0 {
			top := heap.Pop(&hp).(item)
			u, total := top.u, top.cost
			if costs[u] < total {
				continue
			}
			for _, link := range g[u] {
				v, cost := link[0], link[1]
				nextTotal := total + int64(cost)
				if nextTotal <= k && cost >= mid && (costs[v] == -1 || costs[v] > nextTotal) {
					if v+1 == n {
						out = mid
						l = mid + 1
						continue block
					}
					costs[v] = nextTotal
					heap.Push(&hp, item{nextTotal, v})
				}
			}
		}
		r = mid - 1
	}
	return out
}
