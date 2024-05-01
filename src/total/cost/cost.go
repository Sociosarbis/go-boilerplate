package cost

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

const max int = 1e5 + 1

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	l := candidates - 1
	left_hp := h.New(costs[:candidates], func(a, b int) bool {
		return a < b
	})
	r := n - candidates
	if r <= l {
		r = l + 1
	}
	right_hp := h.New(costs[r:], func(a, b int) bool {
		return a < b
	})
	var ret int64
	for k > 0 {
		var left int
		if left_hp.Len() > 0 {
			left = left_hp.Top()
		} else {
			left = max
		}
		var right int
		if right_hp.Len() > 0 {
			right = right_hp.Top()
		} else {
			right = max
		}
		if left <= right {
			ret += int64(left)
			heap.Pop(&left_hp)
			l++
			if l < r {
				heap.Push(&left_hp, costs[l])
			}
		} else {
			ret += int64(right)
			heap.Pop(&right_hp)
			r--
			if l < r {
				heap.Push(&right_hp, costs[r])
			}
		}
		k--
	}
	return ret
}
