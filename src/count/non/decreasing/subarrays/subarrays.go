package subarrays

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	maxHeap := h.New[[2]int]([][2]int{[2]int{0, nums[0]}}, func(a, b [2]int) bool {
		return a[1] > b[1]
	})
	var cost, left, right int
	n := len(nums)
	rightLte := make([][]int, n)
	rightGte := make([]int, n)
	stack := []int{n - 1}
	rightGte[n-1] = n
	for i := n - 2; i >= 0; i-- {
		j := len(stack) - 1
		for ; j >= 0; j-- {
			if nums[i] >= nums[stack[j]] {
				rightLte[i] = append(rightLte[i], stack[j])
			} else {
				break
			}
		}
		stack = append(stack[:j+1], i)
		if len(stack) == 1 {
			rightGte[i] = n
		} else {
			rightGte[i] = stack[len(stack)-2]
		}
	}
	var out int64
	for i := 0; i < n; i++ {
		if i > 0 {
			rlte := rightLte[i-1]
			for _, r := range rlte {
				if r > right {
					break
				}
				// 减少最左侧元素时，减少其右边局部最大元素为了对齐左侧元素的操作次数
				cost -= (min(right+1, rightGte[r]) - r) * (nums[i-1] - nums[r])
			}
		}
		if i > left {
			left = i
		}
		for maxHeap.Len() > 0 && maxHeap.Top()[0] < left {
			heap.Pop(&maxHeap)
		}
		if i > left {
			left = i
		}
		if i > right {
			right = i
			heap.Push(&maxHeap, [2]int{i, nums[i]})
		}
		for right+1 < n {
			top := maxHeap.Top()
			if nums[right+1] < top[1] {
				if cost+top[1]-nums[right+1] <= k {
					cost += top[1] - nums[right+1]
				} else {
					break
				}
			}
			right++
			heap.Push(&maxHeap, [2]int{right, nums[right]})
		}
		out += int64(right - left + 1)
	}
	return out
}
