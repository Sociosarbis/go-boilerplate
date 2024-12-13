package state

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func getFinalState(nums []int, k int, multiplier int) []int {
	n := len(nums)
	arr := make([][2]int, n)
	for i, num := range nums {
		arr[i][0], arr[i][1] = num, i
	}
	hp := h.New(arr, func(a, b [2]int) bool {
		if a[0] < b[0] {
			return true
		} else if a[0] == b[0] {
			return a[1] < b[1]
		}
		return false
	})

	for i := 0; i < k; i++ {
		top := heap.Pop(&hp).([2]int)
		v := top[0] * multiplier
		nums[top[1]] = v
		top[0] = v
		heap.Push(&hp, top)
	}
	return nums
}
