package spending

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

func maxSpending(values [][]int) int64 {
	n := len(values)
	initValues := make([][2]int, 0, n)
	for i, value := range values {
		l := len(value)
		initValues = append(initValues, [2]int{value[l-1], i})
		values[i] = value[:l-1]
	}
	hp := h.New(initValues, func(a, b [2]int) bool {
		return a[0] < b[0]
	})
	var d int64 = 1
	var ret int64
	for hp.Len() != 0 {
		top := heap.Pop(&hp).([2]int)
		ret += d * int64(top[0])
		l := len(values[top[1]])
		if l != 0 {
			heap.Push(&hp, [2]int{values[top[1]][l-1], top[1]})
			values[top[1]] = values[top[1]][:l-1]
		}
		d++
	}
	return ret
}
