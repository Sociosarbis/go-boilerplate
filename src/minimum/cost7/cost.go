package cost7

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func minimumCost(nums []int, k int, dist int) int64 {
	n := len(nums)
	q1 := make([][2]int, 0, k)
	q2 := make([][2]int, 0, k)
	included := make([]bool, n)
	var temp int64
	for i := k - 2; i > 0; i-- {
		q1 = append(q1, [2]int{nums[i], i})
		q2 = append(q2, [2]int{nums[i], i})
		included[i] = true
		temp += int64(nums[i])
	}
	out := int64(nums[0]) + int64(nums[k-1]) + temp
	hp1 := h.New(q1, func(a, b [2]int) bool {
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	hp2 := h.New(q2, func(a, b [2]int) bool {
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	count := hp1.Len()
	for i := k; i < n; i++ {
		minIndex := i - dist
		if minIndex > 0 && included[minIndex-1] {
			included[minIndex-1] = false
			temp -= int64(nums[minIndex-1])
			count--
		}
		heap.Push(&hp2, [2]int{nums[i-1], i - 1})
		for {
			for hp1.Len() > 0 && (count+2 > k || hp1.Top()[1] < minIndex) {
				top := heap.Pop(&hp1).([2]int)
				if included[top[1]] {
					included[top[1]] = false
					temp -= int64(top[0])
					count--
				}
				if top[1] >= minIndex {
					heap.Push(&hp2, top)
				}
			}
			for hp2.Len() > 0 && (count+2 < k || (hp1.Len() == 0 || hp2.Top()[0] < hp1.Top()[0])) {
				top := heap.Pop(&hp2).([2]int)
				if top[1] >= minIndex {
					heap.Push(&hp1, top)
					included[top[1]] = true
					temp += int64(top[0])
					count++
				}
			}
			if count+2 == k {
				break
			}
		}
		if temp+int64(nums[0])+int64(nums[i]) < out {
			out = temp + int64(nums[0]) + int64(nums[i])
		}
	}
	return out
}
