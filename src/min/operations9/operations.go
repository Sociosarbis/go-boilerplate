package operations9

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minOperations(nums []int, k int) int {
	hp := h.New(nums, func(a, b int) bool {
		return a < b
	})
	var ret int
	for {
		a := heap.Pop(&hp).(int)
		if a >= k {
			break
		}
		b := heap.Pop(&hp).(int)
		heap.Push(&hp, a*2+b)
		ret++
	}
	return ret
}
