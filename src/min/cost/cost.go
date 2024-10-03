package cost

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	g := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		g[i] = map[int]int{}
	}

	for _, edge := range edges {
		x, y, time := edge[0], edge[1], edge[2]
		if t, ok := g[x][y]; ok {
			if time < t {
				g[x][y] = time
				g[y][x] = time
			}
		} else {
			g[x][y] = time
			g[y][x] = time
		}
	}

	dp := make([][][2]int, n)
	dp[0] = [][2]int{{passingFees[0], 0}}
	hp := h.New([][3]int{{0, passingFees[0], 0}}, func(a, b [3]int) bool {
		return a[1] < b[1]
	})

	for hp.Len() != 0 {
		top := heap.Pop(&hp).([3]int)
		if top[0]+1 == n {
			if top[2] <= maxTime {
				return top[1]
			}
		} else {
			var found bool
			for _, item := range dp[top[0]] {
				if item[0] == top[1] && item[1] == top[2] {
					found = true
					break
				}
			}
			if found {
			loop:
				for next, t := range g[top[0]] {
					nextTime := t + top[2]
					if nextTime <= maxTime {
						nextCost := top[1] + passingFees[next]
						var l int
						for _, item := range dp[next] {
							if item[0] <= nextCost && item[1] <= nextTime {
								continue loop
							}
						}
						for i, item := range dp[next] {
							if item[0] < nextCost || item[1] < nextTime {
								if i != l {
									dp[next][l] = item
								}
								l++
							}
						}
						if l >= len(dp[next]) {
							dp[next] = append(dp[next], [2]int{nextCost, nextTime})
						} else {
							dp[next][l] = [2]int{nextCost, nextTime}
							dp[next] = dp[next][:l+1]
						}
						heap.Push(&hp, [3]int{next, nextCost, nextTime})
					}
				}
			}
		}
	}

	return -1
}
