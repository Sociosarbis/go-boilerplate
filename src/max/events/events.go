package events

import (
	"container/heap"
	"sort"

	h "github.com/boilerplate/src/heap"
)

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	queue := [][]int{events[0]}
	n := len(events)
	i := 1
	for ; i < n; i++ {
		if events[i][0] == queue[0][0] {
			queue = append(queue, events[i])
		} else {
			break
		}
	}
	hp := h.New(queue, func(a, b []int) bool {
		return a[1] < b[1]
	})
	t := queue[0][0]
	var out int
	for hp.Len() > 0 || i < n {
		if hp.Len() == 0 && t < events[i][0] {
			t = events[i][0]
		}
		for ; i < n; i++ {
			if events[i][0] <= t {
				heap.Push(&hp, events[i])
			} else {
				break
			}
		}
		if hp.Len() > 0 {
			top := heap.Pop(&hp).([]int)
			if top[1] >= t {
				out++
				t++
			}
		} else {
			t++
		}
	}
	return out
}
