package booked

import (
	"container/heap"
	"sort"

	h "github.com/boilerplate/src/heap"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	counter := make([]int, n)
	rooms := make([]int, n)
	for i := range rooms {
		rooms[i] = i
	}
	hp1 := h.New(rooms, func(a, b int) bool {
		return a < b
	})
	usedRooms := [][2]int{}
	hp2 := h.New(usedRooms, func(a, b [2]int) bool {
		return a[0] < b[0]
	})
	var i int
	m := len(meetings)
	var time int
	for i < m {
		start, end := meetings[i][0], meetings[i][1]
		time = max(time, start)
		for hp2.Len() != 0 {
			top := hp2.Top()
			if top[0] <= time {
				heap.Pop(&hp2)
				heap.Push(&hp1, top[1])
			} else {
				break
			}
		}
		if hp1.Len() == 0 {
			top := hp2.Top()
			time = top[0]
		} else {
			top := heap.Pop(&hp1).(int)
			counter[top]++
			heap.Push(&hp2, [2]int{time + end - start, top})
			i++
		}
	}
	var out, temp int
	for i, count := range counter {
		if count > temp {
			out = i
			temp = count
		}
	}
	return out
}
