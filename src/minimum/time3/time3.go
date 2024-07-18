package time3

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := make([]map[int]int, n)
	for i := range g {
		g[i] = map[int]int{}
	}
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if old, ok := g[u][v]; !ok || old > w {
			g[u][v] = w
		}
		if old, ok := g[v][u]; !ok || old > w {
			g[v][u] = w
		}
	}
	hp := h.New([][2]int{{0, 0}}, func(a, b [2]int) bool {
		return a[1] < b[1]
	})
	ret := make([]int, n)
	for hp.Len() != 0 {
		top := heap.Pop(&hp).([2]int)
		if ret[top[0]] > 0 && top[1] > ret[top[0]] {
			continue
		}
		if disappear[top[0]] > top[1] {
			for next, time := range g[top[0]] {
				nextTime := top[1] + time
				if (ret[next] == 0 && next != 0) || ret[next] > nextTime {
					ret[next] = nextTime
					heap.Push(&hp, [2]int{next, nextTime})
				}
			}
		} else {
			ret[top[0]] = -1
		}
	}
	for i := 1; i < n; i++ {
		if ret[i] == 0 {
			ret[i] = -1
		}
	}
	return ret
}
