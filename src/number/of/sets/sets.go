package sets

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}

	for _, r := range roads {
		u, v, w := r[0], r[1], r[2]
		if g[u][v] == 0 || g[u][v] > w {
			g[u][v] = w
			g[v][u] = w
		}
	}

	end := (1 << n) - 1

	ret := 1

	for i := 0; i < end; i++ {
		distance := getMaxDistance(g, i)
		if distance <= maxDistance {
			ret++
		}
	}
	return ret
}

func getMaxDistance(g [][]int, mask int) int {
	n := len(g)
	var ret int
	for i := 0; i < n; i++ {
		var m int
		if mask&(1<<i) == 0 {
			hp := h.New([][2]int{}, func(a, b [2]int) bool {
				return a[1] < b[1]
			})
			heap.Push(&hp, [2]int{i, 0})
			for hp.Len() != 0 {
				top := heap.Pop(&hp).([2]int)
				if m&(1<<top[0]) != 0 {
					continue
				}
				m |= 1 << top[0]
				if ret < top[1] {
					ret = top[1]
				}
				for i := 0; i < n; i++ {
					if (mask|m)&(1<<i) == 0 && g[top[0]][i] != 0 {
						heap.Push(&hp, [2]int{i, top[1] + g[top[0]][i]})
					}
				}
			}
			if (m | mask) != (1<<n)-1 {
				return 1e5 + 1
			}
		}
	}
	return ret
}
