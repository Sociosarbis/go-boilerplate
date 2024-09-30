package manager

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type SeatManager struct {
	hp h.Heap[int]
}

func Constructor(n int) SeatManager {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return SeatManager{
		hp: h.New(arr, func(a, b int) bool {
			return a < b
		}),
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(&this.hp).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.hp, seatNumber)
}
