package tasks

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func assignTasks(servers []int, tasks []int) []int {
	serverIndicesAndWeights := make([][2]int, len(servers))
	for index, weight := range servers {
		serverIndicesAndWeights[index][0], serverIndicesAndWeights[index][1] = index, weight
	}
	freeServers := h.New[[2]int](serverIndicesAndWeights, func(a, b [2]int) bool {
		return a[1] < b[1] || (a[1] == b[1] && a[0] < b[0])
	})
	serverFreeAt := h.New[[3]int]([][3]int{}, func(a, b [3]int) bool {
		return a[2] < b[2]
	})
	out := make([]int, len(tasks))
	var t int
	for i, task := range tasks {
		if t < i {
			t = i
		}
		for serverFreeAt.Len() > 0 {
			top := serverFreeAt.Top()
			if top[2] <= t {
				top = heap.Pop(&serverFreeAt).([3]int)
				heap.Push(&freeServers, [2]int{top[0], top[1]})
			} else {
				break
			}
		}
		if freeServers.Len() == 0 {
			top := heap.Pop(&serverFreeAt).([3]int)
			heap.Push(&freeServers, [2]int{top[0], top[1]})
			t = top[2]
			for serverFreeAt.Len() > 0 {
				top := serverFreeAt.Top()
				if top[2] <= t {
					top = heap.Pop(&serverFreeAt).([3]int)
					heap.Push(&freeServers, [2]int{top[0], top[1]})
				} else {
					break
				}
			}
		}
		top := heap.Pop(&freeServers).([2]int)
		out[i] = top[0]
		heap.Push(&serverFreeAt, [3]int{top[0], top[1], t + task})
	}
	return out
}
