package result

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func maxResult(nums []int, k int) int {
	n := len(nums)
	hp := h.New([][2]int{}, func(a, b [2]int) bool {
		return a[0] > b[0]
	})

	heap.Push(&hp, [2]int{nums[0], 0})

	for i := 1; i < n; i++ {
		num := nums[i]
		maxPrev := 0
		for hp.Len() != 0 {
			item := hp.Top()
			if i-item[1] > k {
				heap.Pop(&hp)
			} else {
				maxPrev = item[0]
				break
			}
		}
		if i+1 == n {
			return num + maxPrev
		} else {
			heap.Push(&hp, [2]int{num + maxPrev, i})
		}
	}
	return nums[0]
}
