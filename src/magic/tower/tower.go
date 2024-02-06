package tower

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func magicTower(nums []int) int {
	var res int64 = 0
	valueHeap := h.New([]int{}, func(a, b int) bool {
		return a < b
	})
	var temp int64 = 1
	ret := 0
	for _, num := range nums {
		if temp+int64(num) <= 0 {
			if valueHeap.Len() != 0 && valueHeap.Top() < num {
				top := heap.Pop(&valueHeap).(int)
				temp -= int64(top)
				heap.Push(&valueHeap, num)
				temp += int64(num)
				res += int64(top)
			} else {
				res += int64(num)
			}
			ret++
		} else {
			if num < 0 {
				heap.Push(&valueHeap, num)
			}
			temp += int64(num)
		}
	}
	if temp+res > 0 {
		return ret
	}
	return -1
}
