package events

import (
	"container/heap"
	"sort"

	h "github.com/boilerplate/src/heap"
)

func maxTwoEvents(events [][]int) int {
	var max int
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	hp := h.New([][]int{}, func(a, b []int) bool {
		return a[1] < b[1]
	})
	var out int
	for _, event := range events {
		for hp.Len() != 0 {
			top := hp.Top()
			if top[1] >= event[0] {
				break
			} else {
				heap.Pop(&hp)
				if top[2] > max {
					max = top[2]
				}
			}
		}

		if event[2]+max > out {
			out = event[2] + max
		}
		heap.Push(&hp, event)
	}
	return out
}
