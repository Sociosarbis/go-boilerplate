package robots

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	times := make([][2]int, n)
	for i, t := range chargeTimes {
		times[i][0], times[i][1] = i, t
	}
	l := 1
	r := n
	var ret int
loop:
	for l <= r {
		mid := (l + r) >> 1
		hp := h.New(times[:mid], func(a, b [2]int) bool {
			return a[1] > b[1]
		})
		var total int64
		for i := 0; i < mid; i++ {
			total += int64(runningCosts[i])
		}
		var temp int64 = int64(hp.Top()[1]) + int64(mid)*int64(total)
		if temp <= budget {
			if mid > ret {
				ret = mid
			}
			l = mid + 1
			continue
		}
		for i := mid; i < n; i++ {
			total += int64(runningCosts[i] - runningCosts[i-mid])
			heap.Push(&hp, times[i])
			for hp.Top()[0] <= i-mid {
				heap.Pop(&hp)
			}
			temp = int64(hp.Top()[1]) + int64(mid)*int64(total)
			if temp <= budget {
				if mid > ret {
					ret = mid
				}
				l = mid + 1
				continue loop
			}
		}
		r = mid - 1
	}
	return ret
}
