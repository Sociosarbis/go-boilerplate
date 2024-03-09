package sum

import (
	"container/heap"
	"sort"

	h "github.com/boilerplate/src/heap"
)

type item struct {
	sum int64
	i   int
}

func kSum(nums []int, k int) int64 {
	var ret int64
	n := len(nums)
	for i, num := range nums {
		if num >= 0 {
			ret += int64(num)
		} else {
			nums[i] = -num
		}
	}
	sort.Ints(nums)
	i := 1
	hp := h.New([]item{{ret, 0}}, func(a, b item) bool {
		return (a.sum-int64(nums[a.i]))-(b.sum-int64(nums[b.i])) > 0
	})
	for i < k {
		top := heap.Pop(&hp).(item)
		temp := top.sum - int64(nums[top.i])
		if i+1 == k {
			return temp
		}
		i++
		if top.i+1 < n {
			heap.Push(&hp, item{temp, top.i + 1})
			heap.Push(&hp, item{top.sum, top.i + 1})
		}
	}
	return ret
}
