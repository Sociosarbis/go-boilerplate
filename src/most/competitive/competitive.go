package competitive

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	if k == n {
		return nums
	}
	items := make([]int, n)

	for i := range items {
		items[i] = i
	}

	hp := h.New(items[:n-k+1], func(a, b int) bool {
		if nums[a] < nums[b] {
			return true
		}
		return nums[a] == nums[b] && a < b
	})

	ret := make([]int, k)
	var min int
	var i int
	for i < k {
		top := heap.Pop(&hp).(int)
		if top >= min {
			min, ret[i] = top, nums[top]
			if i+1 < k {
				heap.Push(&hp, n-k+1+i)
			}
			i++
		}
	}
	return ret
}
