package game

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func numsGame(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	// 将本题化为调整为相同值的最小操作数问题
	// 将所有值调整为中位数时的操作最少
	for i := range nums {
		nums[i] -= i
	}
	left := h.New([]int{}, func(a, b int) bool {
		return a > b
	})

	right := h.New([]int{}, func(a, b int) bool {
		return a < b
	})

	var sumLeft int64 = int64(nums[0])
	left.Push(nums[0])
	var sumRight int64
	for i := 1; i < n; i++ {
		num := nums[i]
		if num <= left.Top() {
			heap.Push(&left, num)
			sumLeft += int64(num)
		} else {
			heap.Push(&right, num)
			sumRight += int64(num)
		}

		if left.Len() == right.Len()+2 {
			temp := int64(heap.Pop(&left).(int))
			sumLeft -= temp
			sumRight += temp
			heap.Push(&right, int(temp))
		} else if left.Len()+1 == right.Len() {
			temp := int64(heap.Pop(&right).(int))
			sumRight -= temp
			sumLeft += temp
			heap.Push(&left, int(temp))
		}

		if left.Len() == right.Len()+1 {
			ret[i] = int((sumRight - sumLeft + int64(left.Top())) % 1000000007)
		} else {
			ret[i] = int((sumRight - sumLeft) % 1000000007)
		}
	}
	return ret
}
