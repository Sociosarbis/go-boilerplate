package paths2

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type item struct {
	i    int
	time int64
}

type optimal struct {
	count int
	time  int64
}

var mask int = 1e9 + 7

func countPaths(n int, roads [][]int) int {
	g := make([]map[int]int, n)
	dp := make([]optimal, n)
	for i := range g {
		g[i] = map[int]int{}
	}

	for i := 1; i < n; i++ {
		dp[i] = optimal{0, -1}
	}

	dp[0] = optimal{1, 0}

	for _, r := range roads {
		a, b, c := r[0], r[1], r[2]
		g[a][b] = c
		g[b][a] = c
	}

	hp := h.New([]item{{0, 0}}, func(a, b item) bool {
		return a.time < b.time
	})

	for {
		top := heap.Pop(&hp).(item)
		if top.time == dp[top.i].time {
			if top.i == n-1 {
				break
			}
			for next, time := range g[top.i] {
				tempTime := int64(time) + dp[top.i].time
				if dp[next].time == -1 || dp[next].time > tempTime {
					dp[next] = optimal{dp[top.i].count, tempTime}
					heap.Push(&hp, item{next, dp[next].time})
				} else if dp[next].time == tempTime {
					dp[next].count = (dp[next].count + dp[top.i].count) % mask
				}
			}
		}
	}
	return dp[n-1].count
}
