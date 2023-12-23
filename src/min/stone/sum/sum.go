package sum

import (
	"container/heap"

	heaputils "github.com/boilerplate/src/heap"
)

func minStoneSum(piles []int, k int) int {
	h := heaputils.New[int](piles,
		func(a, b int) bool {
			return a > b
		},
	)
	for k != 0 {
		pile := heap.Pop(&h).(int)
		if pile&1 == 1 {
			pile = pile/2 + 1
		} else {
			pile = pile / 2
		}
		heap.Push(&h, pile)
		k--
	}
	ret := 0
	for _, num := range piles {
		ret += num
	}
	return ret
}
